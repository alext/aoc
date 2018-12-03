package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

func countLetters(id string) map[rune]int {
	count := make(map[rune]int)
	for _, c := range id {
		count[c] += 1
	}
	return count
}

func checksum(ids []string) int {
	twos, threes := 0, 0
	for _, id := range ids {
		twoInc, threeInc := 0, 0
		for _, count := range countLetters(id) {
			switch count {
			case 2:
				twoInc = 1
			case 3:
				threeInc = 1
			}
		}
		twos += twoInc
		threes += threeInc
	}
	return twos * threes
}

func findMatching(ids []string) string {
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			diffs, diffIndex := 0, 0
			for n := 0; n < len(ids[i]); n++ {
				if ids[i][n] != ids[j][n] {
					diffs++
					diffIndex = n
				}
			}
			if diffs == 1 {
				return ids[i][0:diffIndex] + ids[i][diffIndex+1:]
			}
		}
	}
	return ""
}

func main() {
	ids := make([]string, 0)
	helpers.ScanLines(os.Stdin, func(line string) {
		ids = append(ids, line)
	})
	fmt.Println("Checksum:", checksum(ids))
	fmt.Println("Match:", findMatching(ids))
}
