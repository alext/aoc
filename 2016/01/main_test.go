package main

import "testing"

func TestTurn(t *testing.T) {
	tests := []struct {
		initial  Direction
		turn     string
		expected Direction
	}{
		{
			initial:  North,
			turn:     "L",
			expected: West,
		},
		{
			initial:  North,
			turn:     "R",
			expected: East,
		},
		{
			initial:  East,
			turn:     "L",
			expected: North,
		},
		{
			initial:  East,
			turn:     "R",
			expected: South,
		},
		{
			initial:  South,
			turn:     "L",
			expected: East,
		},
		{
			initial:  South,
			turn:     "R",
			expected: West,
		},
		{
			initial:  West,
			turn:     "L",
			expected: South,
		},
		{
			initial:  West,
			turn:     "R",
			expected: North,
		},
	}

	for _, test := range tests {
		actual := test.initial.Turn(test.turn)
		if actual != test.expected {
			t.Errorf("%s, Turn %s. Want %s, got %s", test.initial, test.turn, test.expected, actual)
		}
	}
}
