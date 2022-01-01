package ui

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
)

const (
	hintTextColor     = "#BB00BB"
	inactiveTextColor = "#999999"
)

type tickMsg struct{}
type errMsg error
type checkWord struct{}
type clearInputMsg struct{}
type showGameSelectorMsg struct{}
type showDifficultySelectorMsg struct{}
type invalidSelectionErrorMsg struct{}

type memorizeItem struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type startGameMsg struct {
	item memorizeItem
}

type textsSuccessfullyLoadedMsg struct {
	Texts []memorizeItem
}

type errorLoadingTextsMsg struct {
	error error
}

type gameDifficulty int

const (
	difficultyLearning gameDifficulty = iota
	difficultyEasy
	difficultyMedium
	difficultyHard
)

type selectDifficultyMsg struct {
	difficulty gameDifficulty
}

func (s gameDifficulty) String() string {
	switch s {
	case difficultyLearning:
		return "Learning"
	case difficultyEasy:
		return "Easy"
	case difficultyMedium:
		return "Medium"
	case difficultyHard:
		return "Hard"
	}
	return "unknown"
}

type gameScreen int

const (
	selectionScreen gameScreen = iota
	gameplayScreen
	selectDifficultyScreen
)

type model struct {
	textInput           textinput.Model
	textComplete        bool
	selectedDifficulty  gameDifficulty
	currentScreen       gameScreen
	memorizeItems       []memorizeItem
	currentMemorizeItem memorizeItem
	uncoveredText       string
	remainingWords      []string
	typed               string
	isTypedWordCorrect  bool
	err                 error
}

func InitialModel(difficulty int) model {
	selectedDifficulty := gameDifficulty(difficulty)
	if selectedDifficulty.String() == "unknown" {
		selectedDifficulty = difficultyMedium
	}
	ti := textinput.NewModel()
	ti.Placeholder = "Start typing"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model{
		textInput:          ti,
		selectedDifficulty: selectedDifficulty,
		currentScreen:      selectionScreen,
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
