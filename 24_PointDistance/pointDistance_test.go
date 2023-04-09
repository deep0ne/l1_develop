package main

import (
	"math"
	"testing"
)

const epsilon = 0.001

func TestPointDistance(t *testing.T) {
	testCases := []struct {
		firstPoint  Point
		secondPoint Point
		expected    float64
	}{
		{
			firstPoint:  NewPoint(10, 5),
			secondPoint: NewPoint(3, 4),
			expected:    math.Sqrt(50),
		},
		{
			firstPoint:  NewPoint(2, 3),
			secondPoint: NewPoint(4, 5),
			expected:    math.Sqrt(8),
		},
		{
			firstPoint:  NewPoint(5, 5),
			secondPoint: NewPoint(4, 5),
			expected:    1,
		},
		{
			firstPoint:  NewPoint(4, 7),
			secondPoint: NewPoint(21, 4),
			expected:    math.Sqrt(298),
		},
		{
			firstPoint:  NewPoint(10, 20),
			secondPoint: NewPoint(10, 10),
			expected:    10,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := PointDistance(tc.firstPoint, tc.secondPoint); math.Abs(output-tc.expected) > epsilon {
			t.Errorf("Point Distance is outputting wrong answer, %f not equal %f", output, tc.expected)
		}
	}

}
