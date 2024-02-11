package main

import (
	"testing"
)

// TestCleanInput verifies that the cleanInput function correctly processes and normalizes input strings.
func TestCleanInput(t *testing.T) {
	// Define test cases with input strings and their expected output after cleaning.
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World", // Test case with mixed case.
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "hello world", // Test case with lower case.
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	// Iterate over each test case, checking if the output from cleanInput matches the expected output.
	for _, testCase := range testCases {
		actual := cleanInput(testCase.input)
		// Verify that the length of the actual output matches the expected output.
		if len(actual) != len(testCase.expected) {
			t.Errorf("Lengths are not equal: expected %v, but got %v", len(testCase.expected), len(actual))
			continue
		}
		// Verify each word in the output matches the expected result.
		for i := range actual {
			actualWord := actual[i]
			expectedWord := testCase.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}
	}
}
