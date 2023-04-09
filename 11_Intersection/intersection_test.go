package main

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		firstSet  map[any]struct{}
		secondSet map[any]struct{}
		expected  []any
	}{
		{
			firstSet: map[any]struct{}{
				"hello": {},
				"world": {},
				"more":  {},
				"words": {},
			},
			secondSet: map[any]struct{}{
				"hello": {},
				"world": {},
			},
			expected: []any{"hello", "world"},
		},
		{
			firstSet: map[any]struct{}{
				1: {},
				2: {},
				3: {},
				4: {},
			},
			secondSet: map[any]struct{}{
				6: {},
				1: {},
			},
			expected: []any{1},
		},
		{
			firstSet: map[any]struct{}{
				"hello": {},
				"world": {},
				"more":  {},
				"words": {},
			},
			secondSet: map[any]struct{}{
				"other": {},
				"word":  {},
			},
			expected: []any{},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		if output := Intersection(tc.firstSet, tc.secondSet); len(output) != len(tc.expected) {
			t.Errorf("Binary search is outputting wrong answer, %s not equal %s", output, tc.expected)
		}
	}

}
