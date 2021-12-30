package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

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
	statusMsg := fmt.Sprintf("%v words remaining", len(m.remainingWords))
	if len(m.remainingWords) == 0 {
		statusMsg = "Complete! Press s to select text."
	}
	typedWord := m.showTypedWord()
	return fmt.Sprintf("%s\n\n%s\nTyped:%s\n%s\n%s",
		m.uncoveredText,
		statusMsg,
		typedWord,
		m.textInput.View(),
		"(esc to cancel)") + "\n"

}

func (m model) View() string {
	docStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF00FF")).
		Height(30).
		Width(60)
	if !m.isPlayingGame {
		return docStyle.Render(m.gameSelector())
	}
	return docStyle.Render(m.gameScreen())
}
