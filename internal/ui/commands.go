package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) startGameCmd(gameIndex int) tea.Cmd {
	newText := m.memorizeItems[gameIndex].Text
	return func() tea.Msg {
		return startGameMsg{text: newText}
	}
}

func (m *model) enterPressedCmd(enteredText string) tea.Cmd {
	enteredNumber, err := strconv.Atoi(enteredText)
	if err != nil || enteredNumber > len(m.memorizeItems)-1 {
		return func() tea.Msg {
			return clearInputMsg{}
		}
	}
	return m.startGameCmd(enteredNumber)
}

func (m *model) loadTextsCmd() tea.Cmd {
	plan, err := ioutil.ReadFile("texts.json")
	if err != nil {
		return func() tea.Msg {
			error := fmt.Errorf("error reading file: %v", err)
			return errorLoadingTextsMsg{error}
		}
	}
	var data []memorizeItem
	err = json.Unmarshal(plan, &data)
	if err != nil {
		return func() tea.Msg {
			error := fmt.Errorf("error parsing JSON: %v", err)
			return errorLoadingTextsMsg{error}
		}
	}
	return func() tea.Msg {
		return textsSuccessfullyLoadedMsg{data}
	}

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
