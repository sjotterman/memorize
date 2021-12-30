package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const totalHeight = 30
const totalWidth = 60

func (m model) gameSelector() string {
	gameSelectorText := "Select a text:\n\n"
	for index, item := range m.memorizeItems {
		gameSelectorText += fmt.Sprintf("%v. %v\n", index, item.title)
	}

	gameSelectorText += m.textInput.View()
	gameSelectorText += "\n(esc to quit)\n"
	return gameSelectorText
}

func (m model) showTypedWord() string {
	if m.typed == "" {
		return ""
	}
	if m.isTypedWordCorrect {
		correctStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
		checkmark := correctStyle.Render("✓")

		return checkmark + " " + m.typed
	}
	incorrectStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	redX := incorrectStyle.Render("✗")
	return redX + " " + m.typed
}

func (m model) gameScreen() string {
	remainingHeight := totalHeight
	statusMsg := fmt.Sprintf("%v words remaining", len(m.remainingWords))
	if len(m.remainingWords) == 0 {
		statusMsg = "Complete! Press s to select text."
	}
	statusMsgHeight := 1
	styledStatusMsg := lipgloss.NewStyle().Height(statusMsgHeight).Render(statusMsg)
	remainingHeight -= statusMsgHeight

	typedWord := m.showTypedWord()
	typedWordHeight := 1
	styledTypedWord := lipgloss.NewStyle().Height(typedWordHeight).Render(typedWord)
	remainingHeight -= typedWordHeight

	textInputHeight := 1
	remainingHeight -= textInputHeight
	styledTextInput := lipgloss.NewStyle().Height(textInputHeight).Render(m.textInput.View())

	helpTextHeight := 1
	remainingHeight -= helpTextHeight
	helpText := "(esc to cancel)"
	styledHelpText := lipgloss.NewStyle().Height(helpTextHeight).Render(helpText)

	uncoveredTextHeight := remainingHeight // needs to be variable
	uncoveredText := lipgloss.NewStyle().Height(uncoveredTextHeight).Render(m.uncoveredText)
	remainingHeight -= uncoveredTextHeight
	return lipgloss.JoinVertical(lipgloss.Left,
		uncoveredText,
		styledStatusMsg,
		styledTypedWord,
		styledTextInput,
		styledHelpText)
}

func (m model) View() string {
	docStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF00FF")).
		Height(totalHeight).
		Width(totalWidth)
	if !m.isPlayingGame {
		return docStyle.Render(m.gameSelector())
	}
	return docStyle.Render(m.gameScreen())
}
