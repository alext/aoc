package main

import "testing"

func TestAddRange(t *testing.T) {
	tests := []struct {
		actual   Blacklist
		addition [2]int64
		expected Blacklist
	}{
		{
			actual:   Blacklist{},
			addition: [2]int64{4, 300},
			expected: Blacklist{{4, 300}},
		},
		{
			actual:   Blacklist{{0, 30}},
			addition: [2]int64{35, 300},
			expected: Blacklist{{0, 30}, {35, 300}},
		},
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{35, 300},
			expected: Blacklist{{35, 300}, {400, 500}},
		},
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{35, 398},
			expected: Blacklist{{35, 398}, {400, 500}},
		},
		// Overlapping ranges
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{350, 450},
			expected: Blacklist{{350, 500}},
		},
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{450, 550},
			expected: Blacklist{{400, 550}},
		},
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{450, 470},
			expected: Blacklist{{400, 500}},
		},
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{350, 400},
			expected: Blacklist{{350, 500}},
		},
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{500, 550},
			expected: Blacklist{{400, 550}},
		},
		// Adjacent ranges
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{501, 550},
			expected: Blacklist{{400, 550}},
		},
		{
			actual:   Blacklist{{400, 500}},
			addition: [2]int64{350, 399},
			expected: Blacklist{{350, 500}},
		},
		// Overlapping end and start
		{
			actual:   Blacklist{{200, 300}, {400, 500}},
			addition: [2]int64{250, 450},
			expected: Blacklist{{200, 500}},
		},
		// Completely overlapping
		{
			actual:   Blacklist{{200, 500}},
			addition: [2]int64{150, 550},
			expected: Blacklist{{150, 550}},
		},
		// Overlapping + adjacent
		{
			actual:   Blacklist{{200, 300}, {400, 500}},
			addition: [2]int64{250, 399},
			expected: Blacklist{{200, 500}},
		},
		// Overlapping multiple ranges
		{
			actual:   Blacklist{{200, 300}, {400, 500}, {600, 700}},
			addition: [2]int64{250, 650},
			expected: Blacklist{{200, 700}},
		},
		// Extremities
		{
			actual:   Blacklist{{100, 200}},
			addition: [2]int64{0, 150},
			expected: Blacklist{{0, 200}},
		},
		{
			actual:   Blacklist{{4285287243, 4288786728}},
			addition: [2]int64{4272117352, MaxAddress},
			expected: Blacklist{{4272117352, MaxAddress}},
		},
	}

	for i, test := range tests {
		test.actual.AddRange(test.addition[0], test.addition[1])

		if test.actual.String() != test.expected.String() {
			t.Errorf("%d: Want %s, Got: %s\n", i, test.expected, test.actual)
			continue
		}
	}
}
