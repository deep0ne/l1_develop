package main

import (
	"reflect"
	"testing"
)

func TestTemperatureGroups(t *testing.T) {
	testCases := []struct {
		temps    []float64
		expected map[int][]float64
	}{
		{
			temps: []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5},
			expected: map[int][]float64{
				-20: {-25.4, -27.0, -21.0},
				10:  {13.0, 19.0, 15.5},
				20:  {24.5},
				30:  {32.5},
			},
		},
		{
			temps: []float64{11.2, 132.4, -12.7, -56.6, -40.3, 19.4, 16.5, -49.8, 24.7, 28.2},
			expected: map[int][]float64{
				-10: {-12.7},
				-40: {-40.3, -49.8},
				-50: {-56.6},
				10:  {11.2, 19.4, 16.5},
				20:  {24.7, 28.2},
				130: {132.4},
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output := TemperatureGroups(tc.temps); !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("Temperature groups not working, %v not equal %v", output, tc.expected)
		}
	}
}
