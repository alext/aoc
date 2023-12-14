package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

var cache = make(map[string]int)

func makeKey(springs string, groups []int) string {
	return fmt.Sprintf("%s%v", springs, groups)
}

func cacheStore(springs string, groups []int, result int) {
	cache[makeKey(springs, groups)] = result
}

func cacheLookup(springs string, groups []int) (int, bool) {
	result, ok := cache[makeKey(springs, groups)]
	return result, ok
}

func CountSolutions(springs string, groups []int, prefix string) int {
	if result, found := cacheLookup(springs, groups); found {
		return result
	}

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
		result := CountSolutions(springs[1:], groups, prefix)
		cacheStore(springs, groups, result)
		return result
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
			result := CountSolutions(springs[groups[0]:], groups[1:], prefix)
			cacheStore(springs, groups, result)
			return result
		}
		if springs[groups[0]] == '#' {
			// # follows this, so not a match
			return 0
		}
		// . or ? follows, so treat as . and continue beyond with rest of groups
		result := CountSolutions(springs[groups[0]+1:], groups[1:], prefix)
		cacheStore(springs, groups, result)
		return result
	case '?':
		// Replace with both . and #
		result := CountSolutions("."+springs[1:], groups, prefix) + CountSolutions("#"+springs[1:], groups, prefix)
		cacheStore(springs, groups, result)
		return result
	default:
		log.Fatalln("Unexpected character", springs[0])
		return 0
	}
}

func main() {

	var springParts []string
	var groupParts [][]int

	helpers.ScanLines(os.Stdin, func(line string) {
		springs, groupsPart, _ := strings.Cut(line, " ")

		var groups []int
		for _, n := range strings.Split(groupsPart, ",") {
			groups = append(groups, helpers.MustAtoi(n))
		}
		springParts = append(springParts, springs)
		groupParts = append(groupParts, groups)
	})

	solutions := 0
	for i, springs := range springParts {
		solutions += CountSolutions(springs, groupParts[i], "")
	}
	fmt.Println("Solutions:", solutions)

	// Part 2
	solutions = 0
	for i := range springParts {
		springs := springParts[i]
		groups := groupParts[i]
		for n := 1; n < 5; n++ {
			springs += "?" + springParts[i]
			groups = append(groups, groupParts[i]...)
		}
		solutions += CountSolutions(springs, groups, "")
	}
	fmt.Println("Expanded solutions:", solutions)
}
