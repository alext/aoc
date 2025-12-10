package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Range struct {
	Start int
	End   int
}

func main() {
	input := helpers.StreamLines(os.Stdin)

	var freshRanges []Range
	for line := range input {
		if line == "" {
			break
		}
		start, end, ok := strings.Cut(line, "-")
		if !ok {
			log.Fatalln("Malformed line", line)
		}
		freshRanges = append(freshRanges, Range{
			Start: helpers.MustAtoi(start),
			End:   helpers.MustAtoi(end),
		})
	}

	freshCount := 0
	for line := range input {
		ingredient := helpers.MustAtoi(line)
		for _, r := range freshRanges {
			if ingredient >= r.Start && ingredient <= r.End {
				freshCount++
				break
			}
		}
	}
	fmt.Println("Fresh count:", freshCount)

	slices.SortFunc(freshRanges, func(a, b Range) int {
		return cmp.Compare(a.Start, b.Start)
	})
	var mergedFreshRanges []Range
	currentRange := freshRanges[0]
	for _, r := range freshRanges[1:] {
		if r.Start <= currentRange.End {
			// overlap
			if r.End > currentRange.End {
				currentRange.End = r.End
			}
			continue
		}
		mergedFreshRanges = append(mergedFreshRanges, currentRange)
		currentRange = r
	}
	mergedFreshRanges = append(mergedFreshRanges, currentRange)

	totalFresh := 0
	for _, r := range mergedFreshRanges {
		totalFresh += r.End - r.Start + 1
	}
	fmt.Println("Total fresh:", totalFresh)
}
