package main

import "testing"

func TestDistance(t *testing.T) {
	tests := []struct {
		ringSize int
		offset   int
		expected int
	}{
		{
			ringSize: 1,
			offset:   1,
			expected: 0,
		},
		{
			ringSize: 3,
			offset:   1,
			expected: 1,
		},
		{
			ringSize: 3,
			offset:   4,
			expected: 2,
		},
		{
			ringSize: 5,
			offset:   1,
			expected: 3,
		},
		{
			ringSize: 5,
			offset:   12,
			expected: 4,
		},
		{
			ringSize: 5,
			offset:   13,
			expected: 3,
		},
	}

	for _, test := range tests {
		actual := calculateDistance(test.ringSize, test.offset)
		if actual != test.expected {
			t.Errorf("Ring: %d, offset: %d: Want %d, got %d", test.ringSize, test.offset, test.expected, actual)
		}
	}
}

func TestDistanceFromAxis(t *testing.T) {
	tests := []struct {
		ringSize int
		offset   int
		expected int
	}{
		{
			ringSize: 1,
			offset:   1,
			expected: 0,
		},
		{
			ringSize: 3,
			offset:   1,
			expected: 0,
		},
		{
			ringSize: 3,
			offset:   4,
			expected: 1,
		},
		{
			ringSize: 5,
			offset:   1,
			expected: 1,
		},
		{
			ringSize: 5,
			offset:   2,
			expected: 0,
		},
		{
			ringSize: 5,
			offset:   12,
			expected: 2,
		},
		{
			ringSize: 5,
			offset:   13,
			expected: 1,
		},
		{
			ringSize: 7,
			offset:   1,
			expected: 2,
		},
		{
			ringSize: 7,
			offset:   3,
			expected: 0,
		},
		{
			ringSize: 7,
			offset:   6,
			expected: 3,
		},
	}

	for _, test := range tests {
		actual := calculateDistanceFromAxis(test.ringSize, test.offset)
		if actual != test.expected {
			t.Errorf("Ring: %d, offset: %d: Want %d, got %d", test.ringSize, test.offset, test.expected, actual)
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

func TestRingSize(t *testing.T) {
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
		actual := ringSize(test.width)
		if actual != test.expected {
			t.Errorf("Width %d: Want %d, got %d", test.width, test.expected, actual)
		}
	}
}
