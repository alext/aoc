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
				Floors: [floors][]string{
					[]string{},
					[]string{},
					[]string{},
					[]string{"something"},
				},
				CurrentFloor: floors - 1,
			},
			Expected: true,
		},
		{
			S: &State{
				Floors: [floors][]string{
					[]string{},
					[]string{"something else"},
					[]string{},
					[]string{"something"},
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

func TestMove(t *testing.T) {

	tests := []struct {
		Initial     *State
		NewFloor    int
		MovingItems []string
		Expected    *State
	}{
		{
			Initial:  &State{},
			NewFloor: 1,
			Expected: &State{CurrentFloor: 1},
		},
		{
			Initial: &State{
				CurrentFloor: 1,
				Floors: [floors][]string{
					[]string{},
					[]string{"alpha", "bravo", "charlie"},
					[]string{"delta", "echo"},
					[]string{"foxtrot"},
				},
			},
			NewFloor:    2,
			MovingItems: []string{"bravo"},
			Expected: &State{
				CurrentFloor: 2,
				Floors: [floors][]string{
					[]string{},
					[]string{"alpha", "charlie"},
					[]string{"bravo", "delta", "echo"},
					[]string{"foxtrot"},
				},
			},
		},
		{
			Initial: &State{
				CurrentFloor: 1,
				Floors: [floors][]string{
					[]string{},
					[]string{"alpha", "bravo", "charlie", "echo"},
					[]string{"delta"},
					[]string{"foxtrot"},
				},
			},
			NewFloor:    0,
			MovingItems: []string{"bravo", "alpha"},
			Expected: &State{
				CurrentFloor: 0,
				Floors: [floors][]string{
					[]string{"alpha", "bravo"},
					[]string{"charlie", "echo"},
					[]string{"delta"},
					[]string{"foxtrot"},
				},
			},
		},
	}
	for _, test := range tests {
		test.Expected.setHash()
		actual := test.Initial.Move(test.NewFloor, test.MovingItems)

		if actual.Hash == 0 {
			t.Fatalf("No hash set for %#v", actual)
		}
		if actual.Hash != test.Expected.Hash {
			t.Errorf("Expected:\n%#v\nGot:\n%#v", test.Expected, actual)
		}
	}
}
