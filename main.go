package main

// A simple CLI program to practice memorizing texts

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type tickMsg struct{}
type errMsg error
type checkWord struct{}
type startGameMsg struct {
	text string
}
type clearInputMsg struct{}
type showGameSelectorMsg struct{}

type memorizeItem struct {
	title string
	text  string
}

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
	textInput      textinput.Model
	textComplete   bool
	isPlayingGame  bool
	uncoveredText  string
	remainingWords []string
	typed          string
	err            error
}

func initialModel() model {
	ti := textinput.NewModel()
	ti.Placeholder = "Start typing"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model{
		textInput:     ti,
		isPlayingGame: false,
		err:           nil,
	}
}

func (m model) startGameCmd(gameIndex int) tea.Cmd {
	log.Println("startGameCmd")
	newText := memorizeItems[gameIndex].text
	return func() tea.Msg {
		return startGameMsg{text: newText}
	}
}

func (m *model) handleStartGameMsg(msg startGameMsg) {
	m.textInput.Reset()
	m.uncoveredText = ""
	m.remainingWords = strings.Split(msg.text, " ")
	m.isPlayingGame = true
	m.textComplete = false
}

func (m model) enterPressedCmd(enteredText string) tea.Cmd {
	enteredNumber, err := strconv.Atoi(enteredText)
	if err != nil || enteredNumber > len(memorizeItems)-1 {
		return func() tea.Msg {
			return clearInputMsg{}
		}
	}
	return m.startGameCmd(enteredNumber)
}

func (m *model) showGameSelectorCmd() tea.Cmd {
	return func() tea.Msg {
		return showGameSelectorMsg{}
	}
}

func (m *model) handleShowGameSelectorMsg() {
	m.textInput.Reset()
	m.isPlayingGame = false
}

func (m *model) handleClearInputMsg() {
	m.textInput.Reset()
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
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
		m.uncoveredText = m.uncoveredText + " " + m.remainingWords[0]
		m.remainingWords = m.remainingWords[1:]
	}
	if len(m.remainingWords) == 0 {
		m.textComplete = true
	}
	m.typed = m.textInput.Value()
	m.textInput.Reset()

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == " " {
			m.checkTypedText()
			return m, nil
		}
		if m.textComplete && msg.String() == "s" {
			return m, m.showGameSelectorCmd()
		}
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			return m, m.enterPressedCmd(m.textInput.Value())
		}

	case errMsg:
		m.err = msg
		return m, nil
	case startGameMsg:
		m.handleStartGameMsg(msg)
		return m, nil
	case clearInputMsg:
		m.handleClearInputMsg()
		return m, nil
	case showGameSelectorMsg:
		m.handleShowGameSelectorMsg()
		return m, nil
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) gameSelector() string {
	gameSelectorText := "Select a text:\n\n"
	for index, item := range memorizeItems {
		gameSelectorText += fmt.Sprintf("%v. %v\n", index, item.title)
	}

	gameSelectorText += m.textInput.View()
	gameSelectorText += "\n(esc to quit)\n"
	return gameSelectorText
}

func (m model) View() string {
	if !m.isPlayingGame {
		return m.gameSelector()
	}
	statusMsg := fmt.Sprintf("%v words remaining", len(m.remainingWords))
	if len(m.remainingWords) == 0 {
		statusMsg = "Complete! Press s to select text."
	}
	return fmt.Sprintf("Start typing: \n>%s\n\n%s\nTyped:%s\n%s\n%s",
		m.uncoveredText,
		statusMsg,
		m.typed,
		m.textInput.View(),
		"(esc to quit)") + "\n"
}
