package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

func SwapPosition(input string, x, y int) string {
	b := []byte(input)
	b[x], b[y] = b[y], b[x]
	return string(b)
}

func SwapLetter(input string, x, y string) string {
	xi := strings.Index(input, x)
	yi := strings.Index(input, y)
	return SwapPosition(input, xi, yi)
}

func Rotate(input string, n int) string {
	b := []byte(input)
	for n < 0 {
		n += len(b)
	}
	if n > len(b) {
		n = n % len(b)
	}
	return string(b[n:]) + string(b[:n])
}

func RotateOnPosition(input string, x string) string {
	xi := strings.Index(input, x)
	if xi >= 4 {
		xi += 1
	}
	return Rotate(input, -(1 + xi))
}

func ReversePositions(input string, x, y int) string {
	b := []byte(input)

	for left, right := x, y; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
	return string(b)
}

func MovePosition(input string, x, y int) string {
	b := []byte(input)
	c := b[x]
	b = append(b[:x], b[x+1:]...)
	b = append(b[:y], append([]byte{c}, b[y:]...)...)
	return string(b)
}

func main() {
	password := "abcdefgh"
	helpers.ScanLines(os.Stdin, func(line string) {
		tokens := strings.Split(line, " ")
		switch tokens[0] + " " + tokens[1] {
		case "swap position":
			password = SwapPosition(password, helpers.MustAtoi(tokens[2]), helpers.MustAtoi(tokens[5]))
		case "swap letter":
			password = SwapLetter(password, tokens[2], tokens[5])
		case "rotate left":
			password = Rotate(password, helpers.MustAtoi(tokens[2]))
		case "rotate right":
			password = Rotate(password, -helpers.MustAtoi(tokens[2]))
		case "rotate based":
			password = RotateOnPosition(password, tokens[6])
		case "reverse positions":
			password = ReversePositions(password, helpers.MustAtoi(tokens[2]), helpers.MustAtoi(tokens[4]))
		case "move position":
			password = MovePosition(password, helpers.MustAtoi(tokens[2]), helpers.MustAtoi(tokens[5]))
		default:
			fmt.Println("Unrecognised input line:", line)
		}
	})
	fmt.Println("Scrambled password:", password)
}
