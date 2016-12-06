package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/alext/aoc/helpers"
)

func main() {
	var countSet []*helpers.LetterCounts
	helpers.ScanLines(os.Stdin, func(line string) {
		for i, r := range line {
			if len(countSet) <= i {
				countSet = append(countSet, new(helpers.LetterCounts))
			}
			countSet[i].Count(r)
		}
	})
	fmt.Print("Message: ")
	for _, counts := range countSet {
		sort.Sort(counts)
		var letter rune
		for _, l := range counts {
			if l.Count == 0 {
				break
			}
			letter = l.Letter
		}
		fmt.Printf("%c", letter)
	}
	fmt.Println("")
}
