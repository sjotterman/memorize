package ui

import (
	"fmt"
	"regexp"

	tea "github.com/charmbracelet/bubbletea"
)

// TODO: place these in the model

// func (m *Model) handleToggleRead(msg tea.Msg, cmds *[]tea.Cmd) {
// 	index := m.list.Index()
// 	newItem := m.items[index]
// 	newItem.isRead = !newItem.isRead
// 	m.items[index] = newItem
// 	*cmds = append(*cmds, m.list.SetItem(index, newItem))
// }

func (m *model) handleShowGameSelectorMsg() {
	m.err = nil
	m.textInput.Reset()
	m.typed = ""
	m.currentScreen = selectionScreen
}

func (m *model) handleShowDifficultySelectorMsg() {
	m.err = nil
	m.textInput.Reset()
	m.typed = ""
	m.currentScreen = selectDifficultyScreen
}

func (m *model) handleSelectDifficultyMsg(msg selectDifficultyMsg) {
	m.textInput.Reset()
	m.err = nil
	m.typed = ""
	m.selectedDifficulty = msg.difficulty
	m.currentScreen = selectionScreen
}

func (m *model) handleInvalidSelectionErrMsg() {
	m.textInput.Reset()
	m.typed = ""
	m.err = fmt.Errorf("Invalid Selection")
}

func (m *model) handleClearInputMsg() {
	m.textInput.Reset()
}

func (m *model) handleErrorLoadingTextsMsg(msg errorLoadingTextsMsg) {
	m.err = msg.error
}

func (m *model) handleStartGameMsg(msg startGameMsg) {
	m.currentMemorizeItem = msg.item
	m.textInput.Reset()
	m.uncoveredText = ""
	re := regexp.MustCompile(`[\s]+`)
	m.remainingWords = re.Split(msg.item.Text, -1)
	m.currentScreen = gameplayScreen
	m.textComplete = false
}

func (m *model) handleTextsLoadedMsg(msg textsSuccessfullyLoadedMsg) {
	m.memorizeItems = msg.Texts
}

func (m *model) handleShowHintMsg() {
	m.showHint = true
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
		if m.currentScreen == selectionScreen && msg.String() == "d" {
			return m, m.handlePressDCmd()
		}
		if m.currentScreen == gameplayScreen && msg.Type == tea.KeyTab {
			return m, m.handleShowHintCmd
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
	case showHintMsg:
		m.handleShowHintMsg()
		return m, nil
	case showGameSelectorMsg:
		m.handleShowGameSelectorMsg()
		return m, nil
	case invalidSelectionErrorMsg:
		m.handleInvalidSelectionErrMsg()
		return m, nil
	case showDifficultySelectorMsg:
		m.handleShowDifficultySelectorMsg()
		return m, nil
	case selectDifficultyMsg:
		m.handleSelectDifficultyMsg(msg)
		return m, nil
	case textsSuccessfullyLoadedMsg:
		m.handleTextsLoadedMsg(msg)
		return m, nil
	case errorLoadingTextsMsg:
		m.handleErrorLoadingTextsMsg(msg)
		return m, nil
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
