package ui

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	textInput          textinput.Model
	textComplete       bool
	isPlayingGame      bool
	selectedDifficulty int
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
		textInput: ti,
		selectedDifficulty: 3,
		isPlayingGame:      false,
		err:                nil,
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
	if len(m.remainingWords) == 0 {
		m.textComplete = true
	}
	m.typed = m.textInput.Value()
	m.textInput.Reset()
}
