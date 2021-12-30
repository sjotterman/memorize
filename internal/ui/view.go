package ui

import "fmt"

func (m model) gameSelector() string {
	gameSelectorText := "Select a text:\n\n"
	for index, item := range m.memorizeItems {
		gameSelectorText += fmt.Sprintf("%v. %v\n", index, item.title)
	}

	gameSelectorText += m.textInput.View()
	gameSelectorText += "\n(esc to quit)\n"
	return gameSelectorText
}

func (m model) gameScreen() string {
	statusMsg := fmt.Sprintf("%v words remaining", len(m.remainingWords))
	if len(m.remainingWords) == 0 {
		statusMsg = "Complete! Press s to select text."
	}
	return fmt.Sprintf("%s\n\n%s\nTyped:%s\n%s\n%s",
		m.uncoveredText,
		statusMsg,
		m.typed,
		m.textInput.View(),
		"(esc to cancel)") + "\n"

}

func (m model) View() string {
	if !m.isPlayingGame {
		return m.gameSelector()
	}
	return m.gameScreen()
}
