package main

import "testing"

func TestCreateSet(t *testing.T) {
	testCases := []struct {
		words          []string
		expectedLength int
	}{
		{
			words:          []string{"cat", "cat", "dog", "cat", "tree"},
			expectedLength: 3,
		},
		{
			words:          []string{"hello", "world", "hello", "world", "hi"},
			expectedLength: 3,
		},
		{
			words:          []string{"hello", "world", "", "", "hi"},
			expectedLength: 4,
		},
		{
			words:          []string{"hello", "world", "world", "world", "world"},
			expectedLength: 2,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := CreateSet(tc.words...); len(output) != tc.expectedLength {
			t.Errorf("Creating set is outputting wrong length, %d not equal %d", len(output), tc.expectedLength)
		}
	}

}
