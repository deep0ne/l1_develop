package main

import (
	"reflect"
	"testing"
)

func TestDeleteFromSlice(t *testing.T) {
	testCases := []struct {
		numbers  []int
		index    int
		expected []int
	}{
		{
			numbers:  []int{7, 3, 5, 11, 23, 1, 0, 6},
			index:    0,
			expected: []int{3, 5, 11, 23, 1, 0, 6},
		},
		{
			numbers:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			index:    10,
			expected: nil,
		},
		{
			numbers:  []int{0, -2, -6, 5, 10, 25},
			index:    3,
			expected: []int{0, -2, -6, 10, 25},
		},
		{
			numbers:  []int{1, 1, 1, 1, 0},
			index:    4,
			expected: []int{1, 1, 1, 1},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := DeleteFromSlice(tc.numbers, tc.index); !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("Delete does not work, %v not equal %v", output, tc.expected)
		}
	}

}
