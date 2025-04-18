package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HeLLo WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  hello",
			expected: []string{"hello", "world", "hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLen := len(actual)
		expectedLen := len(c.expected)

		if actualLen != expectedLen {
			t.Errorf("Different length. Expected: %v, Actual: %v", expectedLen, actualLen)
			t.FailNow()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Different value. Expected: %v, Actual: %v", expectedWord, word)
				t.Fail()
			}
		}
	}
}
