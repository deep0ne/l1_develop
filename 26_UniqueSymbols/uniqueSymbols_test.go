package main

import "testing"

func TestUniqueSymbols(t *testing.T) {
	testCases := []struct {
		word     string
		expected bool
	}{
		{
			word:     "hello",
			expected: false,
		},
		{
			word:     "abcd",
			expected: true,
		},
		{
			word:     "aabcd",
			expected: false,
		},
		{
			word:     "abCdef",
			expected: true,
		},
		{
			word:     "AaBbCcDd",
			expected: false,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := UniqueSymbols(tc.word); output != tc.expected {
			t.Errorf("UniqueSymbols is outputting wrong answer, %v not equal %v", output, tc.expected)
		}
	}

}
