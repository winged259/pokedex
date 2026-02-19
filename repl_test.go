package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  pikachu psYduck",
			expected: []string{"pikachu", "psyduck"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %d, but got %d", len(c.expected), len(actual))
            t.Errorf("Expected words: %v, but got: %v", c.expected, actual)
			continue
		}
		for i, word := range actual {
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word '%s', but got '%s'", expectedWord, word)
			}
		}
	}
}
