package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, testCase := range testCases {
		actual := cleanInput(testCase.input)
		if len(actual) != len(testCase.expected) {
			t.Errorf("Lengths are not equal: expected %v, but got %v", len(testCase.expected), len(actual))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := testCase.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}

		}

	}
}
