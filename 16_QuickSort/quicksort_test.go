package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		numbers  []int
		expected []int
	}{
		{
			numbers:  []int{7, 3, 5, 11, 23, 1, 0, 6},
			expected: []int{0, 1, 3, 5, 6, 7, 11, 23},
		},
		{
			numbers:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			numbers:  []int{0, -2, -6, 5, 10, 25},
			expected: []int{-6, -2, 0, 5, 10, 25},
		},
		{
			numbers:  []int{1, 1, 1, 1, 0},
			expected: []int{0, 1, 1, 1, 1},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := startQuickSort(tc.numbers); !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("QuickSort does not sort, %v not equal %v", output, tc.expected)
		}
	}

}
