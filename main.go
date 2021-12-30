package main

// A simple CLI program to practice memorizing texts

import (
	"fmt"
	"log"
	"regexp"
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

// TODO: remove hardcoded text
const targetText string = "These are the times that try men's souls."

type tickMsg struct{}
type errMsg error
type checkWord struct{}
type startGameMsg struct {
	text string
}

type model struct {
	textInput      textinput.Model
	textComplete   bool
	isTyping       bool
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
	remainingWords := strings.Split(targetText, " ")
	return model{
		textInput:      ti,
		isTyping:       true,
		remainingWords: remainingWords,
		err:            nil,
	}
}

func (m model) startGameCmd() tea.Cmd {
	log.Println("startGameCmd")
	newText := targetText
	return func() tea.Msg {
		return startGameMsg{text: newText}
	}
}

func (m *model) handleStartGameMsg(msg startGameMsg) {
	m.textInput.Reset()
	m.uncoveredText = ""
	m.remainingWords = strings.Split(msg.text, " ")
	m.isTyping = true
	m.textComplete = false
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
		m.isTyping = false
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
		if !m.isTyping && msg.String() == "s" {
			return m, m.startGameCmd()
		}
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	case startGameMsg:
		m.handleStartGameMsg(msg)
		return m, nil
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	statusMsg := fmt.Sprintf("%v words remaining", len(m.remainingWords))
	if len(m.remainingWords) == 0 {
		statusMsg = "Complete! Press s to try again."
	}
	return fmt.Sprintf("Start typing: \n>%s\n\n%s\nTyped:%s\n%s\n%s",
		m.uncoveredText,
		statusMsg,
		m.typed,
		m.textInput.View(),
		"(esc to quit)") + "\n"
}
