package ui

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

const totalHeight = 30
const totalWidth = 65

func (m model) getSelectorTitle(height int) string {
	selectTextMsg := "Select a text:"
	titleText := fmt.Sprintf("Selected difficulty: %v\n\n%v", m.selectedDifficulty, selectTextMsg)
	styledTitleText := lipgloss.NewStyle().Height(height).Render(titleText)
	if m.err != nil {
		styledTitleText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Height(height).Render("Error!")
	}
	return styledTitleText
}

func (m model) getSelectorDisplayText(height int) string {
	mainDisplayText := ""
	for index, item := range m.memorizeItems {
		mainDisplayText += fmt.Sprintf("%v. %v\n", index, item.Title)
	}
	styledListText := lipgloss.NewStyle().Height(height).Render(mainDisplayText)
	return styledListText
}

func (m model) gameSelector() string {
	remainingHeight := totalHeight

	titleTextHeight := 2
	remainingHeight -= titleTextHeight
	styledTitleText := m.getSelectorTitle(titleTextHeight)

	errorMsgHeight := 2
	remainingHeight -= errorMsgHeight
	styledErrorText := m.getErrorText(errorMsgHeight)

	textInputHeight := 1
	remainingHeight -= textInputHeight
	styledTextInput := lipgloss.NewStyle().Height(textInputHeight).Render(m.textInput.View())

	helpTextHeight := 1
	remainingHeight -= helpTextHeight
	helpText := "(esc to quit, d to select difficulty)"
	styledHelpText := lipgloss.NewStyle().Height(helpTextHeight).Render(helpText)

	listTextHeight := remainingHeight

	mainDisplayText := m.getSelectorDisplayText(listTextHeight)
	return lipgloss.JoinVertical(lipgloss.Left,
		styledTitleText,
		mainDisplayText,
		styledErrorText,
		styledTextInput,
		styledHelpText)
}

func (m model) getDifficultySelectorDisplayText(height int) string {
	mainDisplayText := ""
	for i := 0; i <= int(difficultyHard); i++ {
		mainDisplayText += fmt.Sprintf("%v. %v\n", i, gameDifficulty(i).String())
	}
	styledListText := lipgloss.NewStyle().Height(height).Render(mainDisplayText)
	return styledListText
}

func (m model) getErrorText(height int) string {
	errorText := ""
	if m.err != nil {
		errorText = fmt.Sprintf("Error: %v", m.err)
	}
	styledErrorText := lipgloss.NewStyle().Height(height).
		Render(errorText)
	return styledErrorText
}

func (m model) difficultySelector() string {
	remainingHeight := totalHeight

	titleTextHeight := 2
	remainingHeight -= titleTextHeight
	styledTitleText := "Select a difficulty"

	errorMsgHeight := 1
	remainingHeight -= errorMsgHeight
	styledErrorText := m.getErrorText(errorMsgHeight)

	textInputHeight := 1
	remainingHeight -= textInputHeight
	styledTextInput := lipgloss.NewStyle().Height(textInputHeight).Render(m.textInput.View())

	helpTextHeight := 1
	remainingHeight -= helpTextHeight
	helpText := "(esc to quit)"
	styledHelpText := lipgloss.NewStyle().Height(helpTextHeight).Render(helpText)

	listTextHeight := remainingHeight

	mainDisplayText := m.getDifficultySelectorDisplayText(listTextHeight)
	return lipgloss.JoinVertical(lipgloss.Left,
		styledTitleText,
		mainDisplayText,
		styledErrorText,
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

// getWordBlank takes a string and a difficulty level, and returns a partially
// or fully blanked out string to display as a hint
func getWordBlank(word string, difficulty gameDifficulty) string {
	length := utf8.RuneCountInString(word)
	if difficulty == difficultyHard {
		blank := strings.Repeat("_", length)
		return blank
	}
	if difficulty == difficultyLearning {
		return word
	}
	var numLettersToShow int = 0
	if difficulty == difficultyMedium {
		numLettersToShow = 1
	}
	letters := []rune(word)
	if difficulty == difficultyEasy {
		numLettersToShow = int(math.Ceil(float64(length) / 2.0))
	}
	var runes []rune
	for i := 0; i < len(letters); i++ {
		if i < numLettersToShow {
			runes = append(runes, letters[i])
		} else {
			runes = append(runes, '_')
		}
	}
	return string(runes)
}

func (m model) getCoveredWords(difficulty gameDifficulty) string {
	var remainingWordBlanks []string
	for index, word := range m.remainingWords {
		blank := getWordBlank(word, difficulty)
		styledBlank := lipgloss.NewStyle().
			Foreground(lipgloss.Color(inactiveTextColor)).Render(blank)
		if index == 0 {
			nextWord := blank
			if m.showHint {
				nextWord = word
			}
			styledBlank = lipgloss.NewStyle().
				Foreground(lipgloss.Color(hintTextColor)).Render(nextWord)
		}
		remainingWordBlanks = append(remainingWordBlanks, styledBlank)
	}
	coveredWords := strings.Join(remainingWordBlanks, " ")
	return coveredWords
}

func (m model) getGameDisplayText(height int) string {

	coveredWords := m.getCoveredWords(m.selectedDifficulty)
	displayedText := m.uncoveredText
	if m.uncoveredText != "" {
		displayedText += " "
	}

	displayedText += coveredWords
	styledDisplayText := lipgloss.NewStyle().Height(height).Render(displayedText)
	return styledDisplayText
}

func (m model) gameScreen() string {
	remainingHeight := totalHeight

	titleHeight := 2
	remainingHeight -= titleHeight
	styledTitle := lipgloss.NewStyle().Height(titleHeight).Render(m.currentMemorizeItem.Title)

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
	helpText := "(esc to cancel, tab for hint)"
	styledHelpText := lipgloss.NewStyle().Height(helpTextHeight).Render(helpText)

	paragraphHeight := remainingHeight
	styledDisplayText := m.getGameDisplayText(paragraphHeight)
	remainingHeight -= paragraphHeight
	return lipgloss.JoinVertical(lipgloss.Left,
		styledTitle,
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
	if m.currentScreen == selectionScreen {
		return docStyle.Render(m.gameSelector())
	}
	if m.currentScreen == selectDifficultyScreen {
		return docStyle.Render(m.difficultySelector())
	}
	return docStyle.Render(m.gameScreen())
}
