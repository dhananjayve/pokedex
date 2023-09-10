package main

import "testing"

func TestClanInput(t *testing.T) {
	c := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range c {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The length are not equal : %v vs %v",
				len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}

	}
}
