package ui

import (
	"testing"
)

func TestWordBlank(t *testing.T) {

	blankTests := []struct {
		input      string
		difficulty int
		want       string
	}{
		{input: "Hello", difficulty: 4, want: "_____"},
		{input: "I", difficulty: 4, want: "_"},
		{input: "Can't", difficulty: 4, want: "_____"},
		{input: "Hello", difficulty: 3, want: "H____"},
		{input: "I", difficulty: 3, want: "I"},
		{input: "Can't", difficulty: 3, want: "C____"},
		{input: "Hello", difficulty: 2, want: "Hel__"},
		{input: "I", difficulty: 2, want: "I"},
		{input: "Can't", difficulty: 2, want: "Can__"},
		{input: "Hello", difficulty: 1, want: "Hello"},
		{input: "Can't", difficulty: 1, want: "Can't"},
		{input: "I", difficulty: 1, want: "I"},
	}

	for _, tt := range blankTests {
		t.Run(tt.input, func(t *testing.T) {
			got := getWordBlank(tt.input, tt.difficulty)
			if got != tt.want {
				t.Errorf("%v, %v: got %q want %q", tt.input, tt.difficulty, got, tt.want)
			}
		})
	}
}
