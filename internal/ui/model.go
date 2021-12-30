package ui

import (
	"log"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
)

var memorizeItems []memorizeItem = []memorizeItem{
	{
		title: "Opening, The Crisis",
		text:  "These are the times that try men's souls.",
	},
	{
		title: "Hitchhiker's Guide Cover Text",
		text:  "Don't Panic",
	},
}

type model struct {
	textInput          textinput.Model
	textComplete       bool
	isPlayingGame      bool
	memorizeItems      []memorizeItem
	uncoveredText      string
	remainingWords     []string
	typed              string
	isTypedWordCorrect bool
	err                error
}

func InitialModel() model {
	ti := textinput.NewModel()
	ti.Placeholder = "Start typing"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model{
		textInput:     ti,
		isPlayingGame: false,
		err:           nil,
		memorizeItems: memorizeItems,
	}
}

func normalizeWord(input string) string {
	regex := regexp.MustCompile("[[:punct:]]")
	noPunctString := regex.ReplaceAllString(input, "")
	return strings.ToUpper(noPunctString)
}

func (m *model) checkTypedText() {
	targetWord := normalizeWord(m.remainingWords[0])
	typedWord := normalizeWord(m.textInput.Value())
	if targetWord == typedWord {
		m.isTypedWordCorrect = true
		m.uncoveredText = m.uncoveredText + " " + m.remainingWords[0]
		m.remainingWords = m.remainingWords[1:]
	} else {
		m.isTypedWordCorrect = false
	}
	log.Printf("remainingWords: %v\n", m.remainingWords)
	if len(m.remainingWords) == 0 {
		m.textComplete = true
	}
	m.typed = m.textInput.Value()
	m.textInput.Reset()
}
