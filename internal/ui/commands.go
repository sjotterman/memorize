package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) startGameCmd(gameIndex int) tea.Cmd {
	selectedItem := m.memorizeItems[gameIndex]
	return func() tea.Msg {
		return startGameMsg{selectedItem}
	}
}

func (m *model) enterPressedCmd(enteredText string) tea.Cmd {
	enteredNumber, err := strconv.Atoi(enteredText)
	if err != nil {
		return func() tea.Msg {
			return invalidSelectionErrorMsg{}
		}
	}
	if m.currentScreen == selectDifficultyScreen {
		if enteredNumber > int(difficultyHard) {
			return func() tea.Msg {
				return invalidSelectionErrorMsg{}
			}
		}
		selectedDifficulty := gameDifficulty(enteredNumber)
		return func() tea.Msg {
			return selectDifficultyMsg{selectedDifficulty}
		}

	}
	// select text
	if enteredNumber > len(m.memorizeItems)-1 {
		return func() tea.Msg {
			return invalidSelectionErrorMsg{}
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
	if m.currentScreen == gameplayScreen {
		return m.showGameSelectorCmd()
	}
	return tea.Quit
}

func (m *model) handlePressDCmd() tea.Cmd {
	if m.currentScreen == selectionScreen {
		return func() tea.Msg {
			return showDifficultySelectorMsg{}
		}
	}
	return nil
}

func (m *model) showGameSelectorCmd() tea.Cmd {
	return func() tea.Msg {
		return showGameSelectorMsg{}
	}
}
