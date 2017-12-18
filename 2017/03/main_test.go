package main

import "testing"

func TestLocationPosition(t *testing.T) {
	tests := []struct {
		location  int
		expectedX int
		expectedY int
	}{
		{
			location:  1,
			expectedX: 0,
			expectedY: 0,
		},
		{
			location:  2,
			expectedX: 1,
			expectedY: 0,
		},
		{
			location:  5,
			expectedX: -1,
			expectedY: 1,
		},
		{
			location:  10,
			expectedX: 2,
			expectedY: -1,
		},
		{
			location:  13,
			expectedX: 2,
			expectedY: 2,
		},
		{
			location:  17,
			expectedX: -2,
			expectedY: 2,
		},
		{
			location:  22,
			expectedX: -1,
			expectedY: -2,
		},
	}

	for _, test := range tests {
		actualX, actualY := locationPosition(test.location)
		if actualX != test.expectedX || actualY != test.expectedY {
			t.Errorf("Location %d, expected: (%d,%d), got: (%d,%d)", test.location, test.expectedX, test.expectedY, actualX, actualY)
		}
	}
}

func TestRingSizeAndOffset(t *testing.T) {
	tests := []struct {
		location       int
		expectedRing   int
		expectedOffset int
	}{
		{
			location:       1,
			expectedRing:   1,
			expectedOffset: 1,
		},
		{
			location:       2,
			expectedRing:   3,
			expectedOffset: 1,
		},
		{
			location:       9,
			expectedRing:   3,
			expectedOffset: 8,
		},
		{
			location:       10,
			expectedRing:   5,
			expectedOffset: 1,
		},
	}

	for _, test := range tests {
		actualRing, actualOffset := calculateRingSizeAndOffset(test.location)
		if actualRing != test.expectedRing || actualOffset != test.expectedOffset {
			t.Errorf("Location %d: Want (%d,%d), got (%d,%d)", test.location, test.expectedRing, test.expectedOffset, actualRing, actualOffset)
		}
	}
}

func TestRingSizeForWidth(t *testing.T) {
	tests := []struct {
		width    int
		expected int
	}{
		{
			width:    1,
			expected: 1,
		},
		{
			width:    3,
			expected: 8,
		},
		{
			width:    5,
			expected: 16,
		},
	}

	for _, test := range tests {
		actual := ringSizeForWidth(test.width)
		if actual != test.expected {
			t.Errorf("Width %d: Want %d, got %d", test.width, test.expected, actual)
		}
	}
}
