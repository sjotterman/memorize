package ui

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

const totalHeight = 30
const totalWidth = 65

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
	if len(m.remainingWords) <= 0 {
		return ""
	}
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

func (m model) numTypedWords() int {
	re := regexp.MustCompile(`[\s]+`)
	// m.remainingWords = re.Split(msg.text, -1)
	progressWords := re.Split(m.uncoveredText, -1)
	typedWordCount := len(progressWords) - 1
	return typedWordCount
}

func (m model) numTotalWords() int {
	typedWordCount := m.numTypedWords()
	totalWordCount := typedWordCount + len(m.remainingWords)
	return totalWordCount
}

func (m model) fractionComplete() float64 {
	typedWordCount := m.numTypedWords()
	totalWordCount := m.numTotalWords()
	fractionComplete := float64(typedWordCount) / float64(totalWordCount)
	return fractionComplete
}

func (m model) getProgressBar(height int) string {
	fractionComplete := m.fractionComplete()
	progressBarModel := progress.NewModel(progress.WithDefaultScaledGradient())
	if fractionComplete == 1.0 {
		progressBarModel = progress.NewModel(progress.WithSolidFill("#00FF00"))
	}
	progressBarModel.Width = totalWidth
	progressBar := progressBarModel.ViewAs(fractionComplete)
	return progressBar
}

func (m model) getGameStatusMsg(statusMsgHeight int) string {
	statusMsg := fmt.Sprintf("%v/%v words", m.numTypedWords(), m.numTotalWords())
	if len(m.remainingWords) == 0 {
		statusMsg = "Complete! Press s to select another text."
	}
	styledStatusMsg := lipgloss.NewStyle().Height(statusMsgHeight).Render(statusMsg)
	return styledStatusMsg
}

func (m model) gameScreen() string {
	remainingHeight := totalHeight

	statusMsgHeight := 1
	styledStatusMsg := m.getGameStatusMsg(statusMsgHeight)
	remainingHeight -= statusMsgHeight

	typedWord := m.showTypedWord()
	typedWordHeight := 1
	styledTypedWord := lipgloss.NewStyle().Height(typedWordHeight).Render(typedWord)
	remainingHeight -= typedWordHeight

	textInputHeight := 1
	remainingHeight -= textInputHeight
	styledTextInput := lipgloss.NewStyle().Height(textInputHeight).Render(m.textInput.View())

	progressBarHeight := 2
	remainingHeight -= progressBarHeight
	progressBar := m.getProgressBar(progressBarHeight)

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
		progressBar,
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
