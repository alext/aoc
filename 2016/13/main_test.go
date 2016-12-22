package main

import "testing"

func TestIsWall(t *testing.T) {
	seed = 10
	tests := []struct {
		Pos      Position
		Expected bool
	}{
		{
			Pos:      Position{X: 1, Y: 1},
			Expected: false,
		},
		{
			Pos:      Position{X: 1, Y: 0},
			Expected: true,
		},
		{
			Pos:      Position{X: 2, Y: 1},
			Expected: true,
		},
		{
			Pos:      Position{X: 7, Y: 4},
			Expected: false,
		},
	}
	for _, test := range tests {
		actual := test.Pos.IsWall()
		if actual != test.Expected {
			t.Errorf("Expected isWall:%t, got:%t for %v", test.Expected, actual, test.Pos)
		}
	}
}
