package main

import "testing"

func TestComplete(t *testing.T) {
	tests := []struct {
		S        *StateV1
		Expected bool
	}{
		{
			S:        &StateV1{currentFloor: floors - 1},
			Expected: true,
		},
		{
			S:        &StateV1{currentFloor: floors - 2},
			Expected: false,
		},
		{
			S: &StateV1{
				floors: [floors][]string{
					[]string{},
					[]string{},
					[]string{},
					[]string{"something"},
				},
				currentFloor: floors - 1,
			},
			Expected: true,
		},
		{
			S: &StateV1{
				floors: [floors][]string{
					[]string{},
					[]string{"something else"},
					[]string{},
					[]string{"something"},
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

func TestMove(t *testing.T) {

	tests := []struct {
		Initial     *StateV1
		NewFloor    int
		MovingItems []string
		Expected    *StateV1
	}{
		{
			Initial:  &StateV1{},
			NewFloor: 1,
			Expected: &StateV1{currentFloor: 1},
		},
		{
			Initial: &StateV1{
				currentFloor: 1,
				floors: [floors][]string{
					[]string{},
					[]string{"alpha", "bravo", "charlie"},
					[]string{"delta", "echo"},
					[]string{"foxtrot"},
				},
			},
			NewFloor:    2,
			MovingItems: []string{"bravo"},
			Expected: &StateV1{
				currentFloor: 2,
				floors: [floors][]string{
					[]string{},
					[]string{"alpha", "charlie"},
					[]string{"bravo", "delta", "echo"},
					[]string{"foxtrot"},
				},
			},
		},
		{
			Initial: &StateV1{
				currentFloor: 1,
				floors: [floors][]string{
					[]string{},
					[]string{"alpha", "bravo", "charlie", "echo"},
					[]string{"delta"},
					[]string{"foxtrot"},
				},
			},
			NewFloor:    0,
			MovingItems: []string{"bravo", "alpha"},
			Expected: &StateV1{
				currentFloor: 0,
				floors: [floors][]string{
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
		actual := test.Initial.move(test.NewFloor, test.MovingItems)

		if actual.hash == 0 {
			t.Fatalf("No hash set for %#v", actual)
		}
		if actual.hash != test.Expected.hash {
			t.Errorf("Expected:\n%#v\nGot:\n%#v", test.Expected, actual)
		}
	}
}

func TestSafe(t *testing.T) {

	tests := []struct {
		S    *StateV1
		Safe bool
	}{
		{
			S:    &StateV1{},
			Safe: true,
		},
		{
			S: &StateV1{
				floors: [floors][]string{
					[]string{},
					[]string{"alpha microchip", "bravo microchip"},
					[]string{"delta generator"},
					[]string{},
				},
			},
			Safe: true,
		},
		{
			S: &StateV1{
				floors: [floors][]string{
					[]string{"bravo microchip"},
					[]string{},
					[]string{"alpha microchip", "alpha generator", "delta generator"},
					[]string{},
				},
			},
			Safe: true,
		},
		{
			S: &StateV1{
				floors: [floors][]string{
					[]string{},
					[]string{"bravo microchip"},
					[]string{"alpha microchip", "delta generator"},
					[]string{},
				},
			},
			Safe: false,
		},
		{
			S: &StateV1{
				floors: [floors][]string{
					[]string{},
					[]string{"alpha microchip", "bravo microchip", "bravo generator"},
					[]string{"delta generator"},
					[]string{},
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
