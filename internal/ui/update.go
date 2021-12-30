package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

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



func (m *model) handleShowGameSelectorMsg() {
	m.textInput.Reset()
	m.typed = ""
	m.isPlayingGame = false
}

func (m *model) handleClearInputMsg() {
	m.textInput.Reset()
}

func (m *model) handleStartGameMsg(msg startGameMsg) {
	m.textInput.Reset()
	m.uncoveredText = ""
	m.remainingWords = strings.Split(msg.text, " ")
	m.isPlayingGame = true
	m.textComplete = false
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
		case tea.KeyEsc:
			return m, m.escPressedCommand()
		case tea.KeyCtrlC:
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
