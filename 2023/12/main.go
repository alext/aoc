package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

func CountSolutions(springs string, groups []int, prefix string) int {
	//fmt.Printf("%sCountSolutions springs:%s groups:%v\n", prefix, springs, groups)
	if len(springs) == 0 {
		if len(groups) == 0 {
			//fmt.Printf("%s solution\n", prefix)
			return 1
		} else {
			return 0
		}
	}
	prefix = prefix + "  "
	switch springs[0] {
	case '.':
		return CountSolutions(springs[1:], groups, prefix)
	case '#':
		if len(groups) == 0 || len(springs) < groups[0] {
			return 0
		}
		for i := 1; i < groups[0]; i++ {
			if springs[i] == '.' {
				return 0
			}
		}
		// We've found a run of # or ? long enough to fit the next group

		if len(springs) == groups[0] {
			// No more string recurse with the rest (empty) string...
			return CountSolutions(springs[groups[0]:], groups[1:], prefix)
		}
		if springs[groups[0]] == '#' {
			// # follows this, so not a match
			return 0
		}
		// . or ? follows, so treat as . and continue beyond with rest of groups
		return CountSolutions(springs[groups[0]+1:], groups[1:], prefix)
	case '?':
		// Replace with both . and #
		return CountSolutions("."+springs[1:], groups, prefix) + CountSolutions("#"+springs[1:], groups, prefix)
	default:
		log.Fatalln("Unexpected character", springs[0])
		return 0
	}
}

func main() {
	solutions := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		springs, groupsPart, _ := strings.Cut(line, " ")

		var groups []int
		for _, n := range strings.Split(groupsPart, ",") {
			groups = append(groups, helpers.MustAtoi(n))
		}

		solutions += CountSolutions(springs, groups, "")
	})

	fmt.Println("Solutions:", solutions)
}
