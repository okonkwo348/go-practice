package main

import (
	"strings"
	"testing"
)

func TestGenerateArt(t *testing.T) {
	// Mock banner: each character is 8 lines tall
	// Using simple characters to make expected output readable
	banner := map[rune][]string{
		'A': {" A ", "A A", "AAA", "A A", "A A", "A A", "A A", "A A"},
		'B': {"BB ", "B B", "BB ", "B B", "B B", "B B", "B B", "BB "},
		'H': {"H H", "H H", "H H", "HHH", "H H", "H H", "H H", "H H"},
		'i': {" i ", "   ", " i ", " i ", " i ", " i ", " i ", " i "},
	}

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "Single newline",
			input: "\n",
			want:  "\n",
		},
		{
			name:  "Standard word",
			input: "Hi",
			want:  renderExpected(banner, "Hi"),
		},
		{
			name:  "Single newline between letters",
			input: "A\nB",
			want:  renderExpected(banner, "A") + renderExpected(banner, "B"),
		},
		{
			name:  "Double newline between letters",
			input: "A\n\nB",
			want:  renderExpected(banner, "A") + "\n" + renderExpected(banner, "B"),
		},
		{
			name:  "Multiple consecutive newlines",
			input: "\n\n\n",
			want:  "\n\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateArt(tt.input, banner)
			if got != tt.want {
				t.Errorf("GenerateArt() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Helper to construct expected 8-line blocks for testing
func renderExpected(banner map[rune][]string, word string) string {
	var sb strings.Builder
	for i := 0; i < 8; i++ {
		for _, r := range word {
			sb.WriteString(banner[r][i])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
