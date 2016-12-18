package main

import "testing"

func TestComplete2(t *testing.T) {
	tests := []struct {
		S        *StateV2
		Expected bool
	}{
		{
			S:        &StateV2{currentFloor: floors - 1},
			Expected: true,
		},
		{
			S:        &StateV2{currentFloor: floors - 2},
			Expected: false,
		},
		{
			S: &StateV2{
				positions: [][2]uint8{
					[2]uint8{floors - 1, floors - 1},
					[2]uint8{floors - 1, floors - 1},
					[2]uint8{floors - 1, floors - 1},
				},
				currentFloor: floors - 1,
			},
			Expected: true,
		},
		{
			S: &StateV2{
				positions: [][2]uint8{
					[2]uint8{floors - 1, floors - 1},
					[2]uint8{floors - 1, floors - 1},
					[2]uint8{floors - 1, floors - 1},
				},
				currentFloor: 2,
			},
			Expected: false,
		},
		{
			S: &StateV2{
				positions: [][2]uint8{
					[2]uint8{floors - 1, floors - 1},
					[2]uint8{1, floors - 1},
					[2]uint8{floors - 1, floors - 1},
				},
				currentFloor: floors - 1,
			},
			Expected: false,
		},
	}
	for _, test := range tests {
		actual := test.S.Complete()
		if actual != test.Expected {
			t.Errorf("Want %t, got %t for %#v", test.Expected, actual, test.S)
		}
	}
}

func TestSafe2(t *testing.T) {

	tests := []struct {
		S    *StateV2
		Safe bool
	}{
		{
			S:    &StateV2{},
			Safe: true,
		},
		{
			S: &StateV2{
				positions: itemPositions{
					[2]uint8{1, 0},
					[2]uint8{1, 0},
					[2]uint8{3, 2},
				},
			},
			Safe: true,
		},
		{
			S: &StateV2{
				positions: itemPositions{
					[2]uint8{2, 2},
					[2]uint8{0, 1},
					[2]uint8{0, 2},
				},
			},
			Safe: true,
		},
		{
			S: &StateV2{
				positions: itemPositions{
					[2]uint8{2, 0},
					[2]uint8{1, 1},
					[2]uint8{3, 2},
				},
			},
			Safe: false,
		},
		{
			S: &StateV2{
				positions: itemPositions{
					[2]uint8{1, 0},
					[2]uint8{1, 1},
					[2]uint8{2, 3},
				},
			},
			Safe: false,
		},
	}
	for _, test := range tests {
		actualSafe := test.S.Safe()
		if actualSafe != test.Safe {
			t.Errorf("Expected safe:%t, got:%t for %#v", test.Safe, actualSafe, test.S)
		}
	}
}
