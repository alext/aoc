package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

func main() {
	digits := make([]int, 0, 0)
	helpers.ScanRunes(os.Stdin, func(char string) {
		if char == "\n" {
			return
		}

		digit := helpers.MustAtoi(char)
		digits = append(digits, digit)
	})

	total := 0
	offset := len(digits) / 2
	for i, d := range digits {
		j := (i + offset) % len(digits)
		if d == digits[j] {
			total += d
		}
	}
	fmt.Println("Total:", total)
}
