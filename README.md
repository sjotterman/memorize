Memorization app built using Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea) 

# About
Memorize (better name pending) is a simple app that helps you memorize texts by obscuring part or all of a word, letting you fill in the rest.

# How to Use

## Text Selection Screen

- Type in the number of a text, followed by enter, to start playing with that text
- Press d to navigate to the difficulty selection screen
- Press ESC to exit

## Difficulty selection screen
Select a difficulty by typing a number and pressing enter. The available difficulties are:

0. Learning - Remaining words in a text are completely visible
1. Easy - Words are partially obscured - the second half of characters in a word (rounded down) are obscured
2. Medium - Words are mostly obscured - the first letter of a word is visible, but the rest is obscured
3. Hard - Words are completely obscured, leaving only empty spaces to indicate how many characters are in the word

## Game Screen
Play the game by typing words in a text until it has been filled in, pressing space after each typed word.

### Title
The title of the selected text is displayed at the top of the screen.

### Memorization Text
Spaces are shown indicating the number of words and the number of characters in each word. Depending on the selected difficulty, some of the letters will already be filled in.

### Progress
A progress bar displaying how far you are through the current text. Also, a count of current words and total words

### Typed word indicator
Your most recently submitted word will be displayed here, along with a symbol indicating whether it was correct.

### Text input
Type a word and press space to submit. For convenience, punctuation is ignored.

