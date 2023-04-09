package main

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		number   int
		expected int
	}{
		{
			name:     "1 element at the beginning",
			numbers:  []int{5, 8, 10, 12, 14, 16, 21},
			number:   5,
			expected: 0,
		},
		{
			name:     "1 element at the end",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			number:   9,
			expected: 8,
		},
		{
			name:     "1 element at random place",
			numbers:  []int{0, 4, 6, 8, 12, 16, 20, 24, 28, 32, 46, 435},
			number:   8,
			expected: 3,
		},
		{
			name:     "no element",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			number:   12,
			expected: -1,
		},
		{
			name:     "several elements",
			numbers:  []int{1, 2, 3, 4, 5, 5, 6, 7, 9},
			number:   5,
			expected: 4,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := binarySearch(tc.numbers, tc.number); output != tc.expected {
			t.Errorf("Binary search is outputting wrong answer, %d not equal %d", output, tc.expected)
		}
	}

}
