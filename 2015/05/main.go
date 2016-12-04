package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/alext/aoc/2015/helpers"
)

var (
	threeVowels = regexp.MustCompile(`[aeiou].*[aeiou].*[aeiou]`)
	badOnes     = regexp.MustCompile(`ab|cd|pq|xy`)
)

func niceString(str string) bool {
	if !threeVowels.MatchString(str) {
		fmt.Println("No 3 vowels:", str)
		return false
	}
	if badOnes.MatchString(str) {
		fmt.Println("Bad tuples:", str)
		return false
	}
	prev := "."
	doubleLetter := false
	helpers.ScanRunes(strings.NewReader(str), func(c string) {
		if c == prev {
			doubleLetter = true
		}
		prev = c
	})
	if !doubleLetter {
		fmt.Println("No double letter:", str)
		return false
	}
	return true
}

func main() {
	niceCount := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		if niceString(line) {
			niceCount += 1
		}
	})
	fmt.Println("Nice Count:", niceCount)
}
