package main

import "testing"

func TestGenerateRow(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "..^^.",
			Expected: ".^^^^",
		},
		{
			Input:    ".^^^^",
			Expected: "^^..^",
		},
		{
			Input:    ".^^.^.^^^^",
			Expected: "^^^...^..^",
		},
	}

	for _, test := range tests {
		actual := calculateRow([]byte(test.Input))
		if string(actual) != test.Expected {
			t.Errorf("Want %s, got %s for input %s", test.Expected, actual, test.Input)
		}
	}
}
