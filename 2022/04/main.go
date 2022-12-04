package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

func parsePair(in string) (int, int, int, int) {
	first, second, _ := strings.Cut(in, ",")
	a, b, _ := strings.Cut(first, "-")
	c, d, _ := strings.Cut(second, "-")
	return helpers.MustAtoi(a), helpers.MustAtoi(b), helpers.MustAtoi(c), helpers.MustAtoi(d)
}

func main() {
	containCount := 0
	overlapCount := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		xStart, xEnd, yStart, yEnd := parsePair(line)

		if xStart <= yStart && xEnd >= yEnd {
			containCount++
			overlapCount++
			return
		} else if yStart <= xStart && yEnd >= xEnd {
			containCount++
			overlapCount++
			return
		}

		// Check for an overlap that's not contained
		if xStart <= yStart && yStart <= xEnd {
			overlapCount++
		} else if xStart <= yEnd && yEnd <= xEnd {
			overlapCount++
		}
	})

	fmt.Println("Contain count", containCount)
	fmt.Println("Overlap count", overlapCount)
}
