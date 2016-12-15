package main

import "testing"

func TestComplete(t *testing.T) {
	tests := []struct {
		S        *State
		Expected bool
	}{
		{
			S:        &State{CurrentFloor: floors - 1},
			Expected: true,
		},
		{
			S:        &State{CurrentFloor: floors - 2},
			Expected: false,
		},
		{
			S: &State{
				Floors: [floors]Floor{
					Floor{},
					Floor{},
					Floor{},
					Floor{"something"},
				},
				CurrentFloor: floors - 1,
			},
			Expected: true,
		},
		{
			S: &State{
				Floors: [floors]Floor{
					Floor{},
					Floor{"something else"},
					Floor{},
					Floor{"something"},
				},
				CurrentFloor: floors - 1,
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
