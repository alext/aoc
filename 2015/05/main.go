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
)

func niceString(str string) bool {
	if !threeVowels.MatchString(str) {
		fmt.Println("No 3 vowels:", str)
		return false
	}
	if !doubleLetter.MatcherString(str, 0).Matches() {
		fmt.Println("No double letter:", str)
		return false
	}
	if badOnes.MatchString(str) {
		fmt.Println("Bad tuples:", str)
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
