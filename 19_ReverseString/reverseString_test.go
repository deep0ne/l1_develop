package main

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		word     string
		expected string
	}{
		{
			name:     "Ordinary string",
			word:     "hello",
			expected: "olleh",
		},
		{
			name:     "String with spaces",
			word:     "hello world!",
			expected: "!dlrow olleh",
		},
		{
			name:     "String with numbers",
			word:     "hello 1 2 3 4 5 !",
			expected: "! 5 4 3 2 1 olleh",
		},
		{
			name:     "String with chars > 1 byte",
			word:     "главрыба",
			expected: "абырвалг",
		},
		{
			name:     "String with different charbytes",
			word:     "hello мир, как дела? 3 2 1",
			expected: "1 2 3 ?алед как ,рим olleh",
		},
		{
			name:     "Empty string",
			word:     "",
			expected: "",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := ReverseString(tc.word); output != tc.expected {
			t.Errorf("Reverse string is outputting wrong answer, %s not equal %s", output, tc.expected)
		}
	}

}
