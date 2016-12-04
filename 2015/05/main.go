package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/alext/aoc/2015/helpers"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

var (
	threeVowels  = regexp.MustCompile(`[aeiou].*[aeiou].*[aeiou]`)
	doubleLetter = pcre.MustCompile(`([a-z])\1`, 0)
	badOnes      = regexp.MustCompile(`ab|cd|pq|xy`)

	letterPairs    = pcre.MustCompile(`([a-z]{2}).*\1`, 0)
	repeatedLetter = pcre.MustCompile(`([a-z]).\1`, 0)
)

func niceString(str string) bool {
	if !letterPairs.MatcherString(str, 0).Matches() {
		fmt.Println("No letter pairs:", str)
		return false
	}
	if !repeatedLetter.MatcherString(str, 0).Matches() {
		fmt.Println("No repeated letter:", str)
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
