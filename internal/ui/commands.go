package ui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) startGameCmd(gameIndex int) tea.Cmd {
	newText := memorizeItems[gameIndex].text
	return func() tea.Msg {
		return startGameMsg{text: newText}
	}
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

func (m *model) escPressedCommand() tea.Cmd {
	if m.isPlayingGame {
		return m.showGameSelectorCmd()
	}
	return tea.Quit
}

func (m *model) showGameSelectorCmd() tea.Cmd {
	return func() tea.Msg {
		return showGameSelectorMsg{}
	}
}
