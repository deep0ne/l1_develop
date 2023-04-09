package main

import "testing"

func TestReverseSentence(t *testing.T) {
	testCases := []struct {
		name     string
		sentence string
		expected string
	}{
		{
			name:     "Ordinary sentence",
			sentence: "hello world!",
			expected: "world! hello",
		},
		{
			name:     "Test from L1",
			sentence: "snow dog sun",
			expected: "sun dog snow",
		},
		{
			name:     "numbers",
			sentence: "1, 2, 3 4 5 6 7",
			expected: "7 6 5 4 3 2, 1,",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := ReverseSentence(tc.sentence); output != tc.expected {
			t.Errorf("Reverse sentence is outputting wrong answer, %s not equal %s", output, tc.expected)
		}
	}

}
