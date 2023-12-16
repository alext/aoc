package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Grid [][]string

func (g Grid) String() string {
	var b = strings.Builder{}
	for i, line := range g {
		b.WriteString(strings.Join(line, ""))
		if i < len(g)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func (g Grid) MoveNorth() {
	if len(g) == 0 {
		return
	}
	colNextClearRow := make([]int, len(g[0]))

	for row, line := range g {
		for col, thing := range line {
			if thing == "O" && colNextClearRow[col] < row {
				g[colNextClearRow[col]][col] = "O"
				g[row][col] = "."
				colNextClearRow[col]++
			}
			if g[row][col] != "." {
				colNextClearRow[col] = row + 1
			}
		}
		//fmt.Printf("After row %d, %v\n", row, colNextClearRow)
	}
}

func (g Grid) TotalLoad() int {
	total := 0
	for row, line := range g {
		rowLoad := len(g) - row
		for _, thing := range line {
			if thing == "O" {
				total += rowLoad
			}
		}
	}
	return total
}

func main() {
	grid := Grid(helpers.ScanGrid(helpers.StreamLines(os.Stdin), ""))

	fmt.Println("Before:")
	fmt.Println(grid)

	grid.MoveNorth()
	fmt.Println("After:")
	fmt.Println(grid)
	fmt.Println("Total load:", grid.TotalLoad())
}
