package hexgrid

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var testCases = []struct {
		hexA Hexagon
		hexB Hexagon
		expected Hexagon
	}{
		{NewHexQR(2, -4), NewHexQR(-1, 2), NewHexQR(1, -2)},
	}

	for _, tc := range testCases {
		actual := Add(tc.hexA, tc.hexB)

		if actual != tc.expected {
			t.Error("Expected: ", tc.expected, "got: ", actual)
		}

	}
}

func TestScale(t *testing.T) {
	var testCases = []struct {
		hex Hexagon
		radius int64
		expected Hexagon
	} {
		{NewHexQR(2, 1), 2, NewHexXYZ(4, 2, -6)},
	}

	for _, tc := range testCases {
		actual := Scale(tc.hex, tc.radius)

		if(actual != tc.expected) {
			t.Error("Expected: ", tc.expected, "got: ", actual)
		}
	}
}
