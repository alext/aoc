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
		fmt.Printf("%c", counts[0].Letter)
	}
	fmt.Println("")
}
