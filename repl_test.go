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
			input:    "  hello     world i'm  gonna be a developer  !  ",
			expected: []string{"hello", "world", "i'm", "gonna", "be", "a", "developer", "!"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "\thello\tworld\nnewline\ttest ",
			expected: []string{"hello", "world", "newline", "test"},
		},
		{
			input:    "punctuation! is, tricky.",
			expected: []string{"punctuation!", "is,", "tricky."},
		},
		{
			input:    "multiple\n\n\nnewlines\nbetween\nwords",
			expected: []string{"multiple", "newlines", "between", "words"},
		},
		{
			input:    " mix OF     CAPS   and lowercase ",
			expected: []string{"mix", "of", "caps", "and", "lowercase"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Slice lengths do not match! (actual = %d, expected = %d)", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match! (actual = %s, expected = %s)", word, expectedWord)
			}
		}
	}
}
