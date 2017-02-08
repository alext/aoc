package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

func handleWrap(n, max int) int {
	for n < 0 {
		n += max
	}
	return n % max
}

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
	n = handleWrap(n, len(b))
	return string(b[n:]) + string(b[:n])
}

func RotateOnPosition(input string, x string) string {
	xi := strings.Index(input, x)
	if xi >= 4 {
		xi += 1
	}
	return Rotate(input, -(1 + xi))
}

func ReverseRotateOnPosition(input string, x string) string {
	final := strings.Index(input, x)
	initial := handleWrap((final-1)/2, len(input))
	if final%2 == 0 {
		// was >= 4
		initial = handleWrap(final-2, len(input))/2 + 4
	}

	return Rotate(input, final-initial)
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

func processLine(password, line string, reverse bool) string {
	tokens := strings.Split(line, " ")
	switch tokens[0] + " " + tokens[1] {
	case "swap position":
		return SwapPosition(password, helpers.MustAtoi(tokens[2]), helpers.MustAtoi(tokens[5]))
	case "swap letter":
		return SwapLetter(password, tokens[2], tokens[5])
	case "rotate left":
		if reverse {
			return Rotate(password, -helpers.MustAtoi(tokens[2]))
		} else {
			return Rotate(password, helpers.MustAtoi(tokens[2]))
		}
	case "rotate right":
		if reverse {
			return Rotate(password, helpers.MustAtoi(tokens[2]))
		} else {
			return Rotate(password, -helpers.MustAtoi(tokens[2]))
		}
	case "rotate based":
		if reverse {
			return ReverseRotateOnPosition(password, tokens[6])
		} else {
			return RotateOnPosition(password, tokens[6])
		}
	case "reverse positions":
		return ReversePositions(password, helpers.MustAtoi(tokens[2]), helpers.MustAtoi(tokens[4]))
	case "move position":
		if reverse {
			return MovePosition(password, helpers.MustAtoi(tokens[5]), helpers.MustAtoi(tokens[2]))
		} else {
			return MovePosition(password, helpers.MustAtoi(tokens[2]), helpers.MustAtoi(tokens[5]))
		}
	default:
		fmt.Println("Unrecognised input line:", line)
		return password
	}
}

func main() {
	input := flag.String("input", "", "the password to scramble or unscramble")
	reverse := flag.Bool("reverse", false, "whether to apply the scrambling in reverse")
	flag.Parse()
	if *input == "" {
		log.Fatal("No input provided")
	}
	if *reverse && len(*input) != 8 {
		log.Fatal("Reverse only works for 8 char input")
	}

	var lines []string
	helpers.ScanLines(os.Stdin, func(line string) {
		lines = append(lines, line)
	})
	password := *input
	if *reverse {
		for i := len(lines) - 1; i >= 0; i-- {
			password = processLine(password, lines[i], *reverse)
		}
	} else {
		for _, line := range lines {
			password = processLine(password, line, *reverse)
		}
	}
	fmt.Println("Scrambled password:", password)
}
