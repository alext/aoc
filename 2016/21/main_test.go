package main

import "testing"

func TestSwapPosition(t *testing.T) {
	tests := []struct {
		input  string
		x      int
		y      int
		output string
	}{
		{
			input:  "abcde",
			x:      1,
			y:      3,
			output: "adcbe",
		},
		{
			input:  "abcde",
			x:      1,
			y:      1,
			output: "abcde",
		},
	}
	for i, test := range tests {
		actual := SwapPosition(test.input, test.x, test.y)
		if actual != test.output {
			t.Errorf("%d: Want %s, Got: %s (input: %s, x:%d, y:%d)\n", i, test.output, actual, test.input, test.x, test.y)
		}
	}
}

func TestSwapLetter(t *testing.T) {
	tests := []struct {
		input  string
		x      string
		y      string
		output string
	}{
		{
			input:  "abcde",
			x:      "b",
			y:      "e",
			output: "aecdb",
		},
	}
	for i, test := range tests {
		actual := SwapLetter(test.input, test.x, test.y)
		if actual != test.output {
			t.Errorf("%d: Want %s, Got: %s (input: %s, x:%s, y:%s)\n", i, test.output, actual, test.input, test.x, test.y)
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		input  string
		n      int
		output string
	}{
		{
			input:  "abcde",
			n:      1,
			output: "bcdea",
		},
		{
			input:  "abcde",
			n:      -2,
			output: "deabc",
		},
		{
			input:  "abcde",
			n:      6,
			output: "bcdea",
		},
		{
			input:  "abcde",
			n:      -6,
			output: "eabcd",
		},
	}
	for i, test := range tests {
		actual := Rotate(test.input, test.n)
		if actual != test.output {
			t.Errorf("%d: Want %s, Got: %s (input: %s, n:%d)\n", i, test.output, actual, test.input, test.n)
		}
	}
}

func TestRotateOnPosition(t *testing.T) {
	tests := []struct {
		input  string
		x      string
		output string
	}{
		{
			input:  "abcde",
			x:      "a",
			output: "eabcd",
		},
		{
			input:  "abcde",
			x:      "b",
			output: "deabc",
		},
		{
			input:  "abcdefgh",
			x:      "e",
			output: "cdefghab",
		},
	}
	for i, test := range tests {
		actual := RotateOnPosition(test.input, test.x)
		if actual != test.output {
			t.Errorf("%d: Want %s, Got: %s (input: %s, x:%s)\n", i, test.output, actual, test.input, test.x)
		}
	}
}

func TestReverseRotateOnPosition(t *testing.T) {
	input := "abcdefgh"
	for i := 0; i < 8; i++ {
		char := string(input[i])
		rotated := RotateOnPosition(input, char)
		result := ReverseRotateOnPosition(rotated, char)
		if result != input {
			t.Errorf("Want %s, Got %s (char %s)", input, result, char)
		}
	}
}

func TestReversePositions(t *testing.T) {
	tests := []struct {
		input  string
		x      int
		y      int
		output string
	}{
		{
			input:  "abcde",
			x:      1,
			y:      3,
			output: "adcbe",
		},
		{
			input:  "abcdefghijk",
			x:      0,
			y:      5,
			output: "fedcbaghijk",
		},
	}
	for i, test := range tests {
		actual := ReversePositions(test.input, test.x, test.y)
		if actual != test.output {
			t.Errorf("%d: Want %s, Got: %s (input: %s, x:%d, y:%d)\n", i, test.output, actual, test.input, test.x, test.y)
		}
	}
}

func TestMovePosition(t *testing.T) {
	tests := []struct {
		input  string
		x      int
		y      int
		output string
	}{
		{
			input:  "abcde",
			x:      1,
			y:      3,
			output: "acdbe",
		},
		{
			input:  "abcdefghijk",
			x:      5,
			y:      2,
			output: "abfcdeghijk",
		},
		{
			input:  "abcde",
			x:      2,
			y:      2,
			output: "abcde",
		},
	}
	for i, test := range tests {
		actual := MovePosition(test.input, test.x, test.y)
		if actual != test.output {
			t.Errorf("%d: Want %s, Got: %s (input: %s, x:%d, y:%d)\n", i, test.output, actual, test.input, test.x, test.y)
		}
	}
}
