package ui

import (
	"testing"
)

func TestWordBlank(t *testing.T) {

	blankTests := []struct {
		input      string
		difficulty gameDifficulty
		want       string
	}{
		{input: "Hello", difficulty: difficultyHard, want: "_____"},
		{input: "I", difficulty: difficultyHard, want: "_"},
		{input: "Can't", difficulty: difficultyHard, want: "_____"},
		{input: "Hello", difficulty: difficultyMedium, want: "H____"},
		{input: "I", difficulty: difficultyMedium, want: "I"},
		{input: "Can't", difficulty: difficultyMedium, want: "C____"},
		{input: "Hello", difficulty: difficultyEasy, want: "Hel__"},
		{input: "I", difficulty: difficultyEasy, want: "I"},
		{input: "Can't", difficulty: difficultyEasy, want: "Can__"},
		{input: "Hello", difficulty: difficultyLearning, want: "Hello"},
		{input: "Can't", difficulty: difficultyLearning, want: "Can't"},
		{input: "I", difficulty: difficultyLearning, want: "I"},
	}

	for _, tt := range blankTests {
		t.Run(tt.input, func(t *testing.T) {
			got := getWordBlank(tt.input, tt.difficulty)
			if got != tt.want {
				t.Errorf("%v, %v: got %q want %q", tt.input, tt.difficulty.String(), got, tt.want)
			}
		})
	}
}
