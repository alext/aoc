package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestLookupRange(t *testing.T) {
	m := &Map{Ranges: []MapRange{
		{src: 5, dest: 105, length: 5},
	}}

	tests := []struct {
		Input    Range
		Expected []Range
	}{
		// Inside range
		{
			Input:    Range{Start: 6, End: 9},
			Expected: []Range{{Start: 106, End: 109}},
		},
		// No overlap
		{
			Input:    Range{Start: 1, End: 4},
			Expected: []Range{{Start: 1, End: 4}},
		},
		// Overlap end
		{
			Input:    Range{Start: 7, End: 12},
			Expected: []Range{{Start: 107, End: 109}, {Start: 10, End: 12}},
		},
		// Overlap start
		{
			Input:    Range{Start: 3, End: 7},
			Expected: []Range{{Start: 105, End: 107}, {Start: 3, End: 4}},
		},
		// Overlap both
		{
			Input:    Range{Start: 3, End: 11},
			Expected: []Range{{Start: 105, End: 109}, {Start: 3, End: 4}, {Start: 10, End: 11}},
		},
	}
	for _, test := range tests {
		fmt.Println("Test start")
		actual := m.LookupRange(test.Input)

		if !slices.Equal(actual, test.Expected) {
			t.Errorf("Expected: %v, got: %v. Input %v", test.Expected, actual, test.Input)
		}
	}
}
