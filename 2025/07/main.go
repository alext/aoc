package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

func countSplits(grid [][]string) int {
	splits := 0

	for r := 0; r < len(grid)-1; r++ {
		for c := 0; c < len(grid[r]); c++ {
			switch grid[r][c] {
			case "S":
				grid[r+1][c] = "|"
			case "|":
				if grid[r+1][c] == "^" {
					splits++
					grid[r+1][c+1] = "|"
					grid[r+1][c-1] = "|"
				} else {
					grid[r+1][c] = "|"
				}
			}
		}
	}

	return splits
}

func countTimelines(grid [][]string) int {
	var currentRow, nextRow []int
	for r := 0; r < len(grid)-1; r++ {
		nextRow = make([]int, len(grid[r]))
		for c := 0; c < len(grid[r]); c++ {
			switch grid[r][c] {
			case "S":
				nextRow[c] = 1
			case "|":
				if grid[r+1][c] == "^" {
					nextRow[c+1] += currentRow[c]
					nextRow[c-1] += currentRow[c]
				} else {
					nextRow[c] += currentRow[c]
				}
			}
		}
		currentRow = nextRow
	}

	timelines := 0
	for _, n := range currentRow {
		timelines += n
	}
	return timelines
}

func main() {
	grid := helpers.ScanGrid(os.Stdin, "")

	fmt.Println("Total splits:", countSplits(grid))
	fmt.Println("Total timelines:", countTimelines(grid))
}
