package main

import "testing"

func TestDial(t *testing.T) {
	testCases := []struct {
		initial  int
		move     int
		expected int
	}{
		{50, 30, 0},
		{0, 30, 0},
		{0, 99, 0},
		{0, 100, 1},
		{0, 101, 1},
		{0, 200, 2},
		{0, -30, 0},
		{0, -99, 0},
		{0, -100, 1},
		{0, -101, 1},
		{0, -199, 1},
		{0, -200, 2},
		{0, -201, 2},
	}
	for _, tc := range testCases {
		d := Dial{Pos: tc.initial}
		d.Move(tc.move)
		if d.PassZero != tc.expected {
			t.Errorf("Initial:%d, Move:%d; want: %d, got %d (finalPos: %d)", tc.initial, tc.move, tc.expected, d.PassZero, d.Pos)

		}
	}
}
