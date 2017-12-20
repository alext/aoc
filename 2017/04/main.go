package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/alext/aoc/helpers"
)

func main() {
	valid1Count := 0
	valid2Count := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		if validPassword1(line) {
			valid1Count++
		}
		if validPassword2(line) {
			valid2Count++
		}
	})
	fmt.Println("Valid passwords:", valid1Count)
	fmt.Println("Valid passwords by second policy:", valid2Count)
}

func validPassword1(password string) bool {
	seenWords := make(map[string]bool)
	for _, word := range strings.Split(password, " ") {
		if seenWords[word] {
			return false
		}
		seenWords[word] = true
	}
	return true
}

func validPassword2(password string) bool {
	seenWords := make(map[string]bool)
	for _, word := range strings.Split(password, " ") {
		sortedWord := sortString(word)
		if seenWords[sortedWord] {
			return false
		}
		seenWords[sortedWord] = true
	}
	return true
}

func sortString(input string) string {
	runes := []rune(input)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
