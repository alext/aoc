package main

import "testing"

func TestRoundPos(t *testing.T) {
	tests := []struct {
		Input    Pos
		Expected Pos
	}{
		// Inside range
		{
			Input:    Pos{X: 10, Y: 10},
			Expected: Pos{X: 10, Y: 10},
		},
		{
			Input:    Pos{X: 0, Y: 0},
			Expected: Pos{X: 0, Y: 0},
		},
		{
			Input:    Pos{X: 130, Y: 130},
			Expected: Pos{X: 130, Y: 130},
		},
		// Larger
		{
			Input:    Pos{X: 131, Y: 131},
			Expected: Pos{X: 0, Y: 0},
		},
		{
			Input:    Pos{X: 262, Y: 132},
			Expected: Pos{X: 0, Y: 1},
		},
		// Negative
		{
			Input:    Pos{X: -1, Y: -1},
			Expected: Pos{X: 130, Y: 130},
		},
		{
			Input:    Pos{X: -131, Y: -131},
			Expected: Pos{X: 0, Y: 0},
		},
		{
			Input:    Pos{X: -132, Y: -132},
			Expected: Pos{X: 130, Y: 130},
		},
	}
	for _, test := range tests {
		actual := roundPos(test.Input, 131)

		if actual != test.Expected {
			t.Errorf("Expected: %s, got: %s. Input: %s", test.Expected, actual, test.Input)
		}
	}
}
