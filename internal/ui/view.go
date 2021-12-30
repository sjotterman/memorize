package ui

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

const totalHeight = 30
const totalWidth = 60

func (m model) gameSelector() string {
	remainingHeight := totalHeight

	titleTextHeight := 2
	remainingHeight -= titleTextHeight
	titleText := "Select a text"
	styledTitleText := lipgloss.NewStyle().Height(titleTextHeight).Render(titleText)

	textInputHeight := 1
	remainingHeight -= textInputHeight
	styledTextInput := lipgloss.NewStyle().Height(textInputHeight).Render(m.textInput.View())

	helpTextHeight := 1
	remainingHeight -= helpTextHeight
	helpText := "(esc to quit)"
	styledHelpText := lipgloss.NewStyle().Height(helpTextHeight).Render(helpText)

	listText := ""
	for index, item := range m.memorizeItems {
		listText += fmt.Sprintf("%v. %v\n", index, item.title)
	}
	listTextHeight := remainingHeight
	styledListText := lipgloss.NewStyle().Height(listTextHeight).Render(listText)

	return lipgloss.JoinVertical(lipgloss.Left,
		styledTitleText,
		styledListText,
		styledTextInput,
		styledHelpText)
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

	var remainingWordBlanks []string
	for _, word := range m.remainingWords {
		length := utf8.RuneCountInString(word)
		blank := strings.Repeat("_", length)
		remainingWordBlanks = append(remainingWordBlanks, blank)
	}
	coveredWords := strings.Join(remainingWordBlanks, " ")
	paragraphHeight := remainingHeight
	displayedText := m.uncoveredText
	if m.uncoveredText != "" {
		displayedText += " "
	}

	displayedText += lipgloss.NewStyle().Foreground(lipgloss.Color("#999999")).Render(coveredWords)
	styledDisplayText := lipgloss.NewStyle().Height(paragraphHeight).Render(displayedText)
	remainingHeight -= paragraphHeight
	return lipgloss.JoinVertical(lipgloss.Left,
		styledDisplayText,
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
