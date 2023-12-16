package main

import (
	"fmt"
	"os"
	"slices"
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

func (g Grid) Clone() Grid {
	clone := make(Grid, 0, len(g))
	for _, row := range g {
		clone = append(clone, slices.Clone(row))
	}
	return clone
}

func (g Grid) moveVertical(north bool) {
	colNextClearRow := make([]int, len(g[0]))
	var startRow, increment int
	if north {
		startRow = 0
		increment = 1
	} else {
		for i := range colNextClearRow {
			colNextClearRow[i] = len(g) - 1
		}
		startRow = len(g) - 1
		increment = -1
	}
	for row := startRow; row >= 0 && row < len(g); row += increment {
		for col := 0; col < len(g[row]); col++ {
			if g[row][col] == "O" && colNextClearRow[col] != row {
				g[colNextClearRow[col]][col] = "O"
				g[row][col] = "."
				colNextClearRow[col] += increment
			}
			if g[row][col] != "." {
				colNextClearRow[col] = row + increment
			}
		}
	}
}

func (g Grid) moveHorizontal(west bool) {
	rowNextClearCol := make([]int, len(g))
	var startCol, increment int
	if west {
		startCol = 0
		increment = 1
	} else {
		for i := range rowNextClearCol {
			rowNextClearCol[i] = len(g[0]) - 1
		}
		startCol = len(g[0]) - 1
		increment = -1
	}
	for col := startCol; col >= 0 && col < len(g[0]); col += increment {
		for row := 0; row < len(g); row++ {
			if g[row][col] == "O" && rowNextClearCol[row] != col {
				g[row][rowNextClearCol[row]] = "O"
				g[row][col] = "."
				rowNextClearCol[row] += increment
			}
			if g[row][col] != "." {
				rowNextClearCol[row] = col + increment
			}
		}
	}
}

func (g Grid) MoveNorth() { g.moveVertical(true) }
func (g Grid) MoveSouth() { g.moveVertical(false) }
func (g Grid) MoveWest()  { g.moveHorizontal(true) }
func (g Grid) MoveEast()  { g.moveHorizontal(false) }

func (g Grid) SpinCycle() {
	g.MoveNorth()
	g.MoveWest()
	g.MoveSouth()
	g.MoveEast()
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

var seenArrangements = make(map[string]int)

func checkArrangementSeen(g Grid, thisIteration int) (bool, int) {
	key := g.String()
	if i, found := seenArrangements[key]; found {
		return true, i
	}
	seenArrangements[key] = thisIteration
	return false, 0
}

func main() {
	grid := Grid(helpers.ScanGrid(helpers.StreamLines(os.Stdin), ""))

	fmt.Println("Before:")
	fmt.Println(grid)

	grid2 := grid.Clone()

	grid.MoveNorth()
	fmt.Println("After:")
	fmt.Println(grid)
	fmt.Println("Total load:", grid.TotalLoad())

	grid = grid2

	fmt.Println("\nPart 2")
	fmt.Println(grid)

	const cycleCount = 1000000000
	//const cycleCount = 22

	var loopStart, loopEnd int
	for i := 1; i <= cycleCount; i++ {
		grid.SpinCycle()

		//fmt.Printf("After cycle %d:\n%s\n", i, grid)

		seen, seenAt := checkArrangementSeen(grid, i)
		if seen {
			fmt.Println("Arrangement previously seen at cycle", seenAt)
			loopStart = seenAt
			loopEnd = i
			break
		}
	}

	fmt.Printf("Loops between cycle %d and %d\n", loopStart, loopEnd)
	loopLength := loopEnd - loopStart
	fmt.Println("Loop length:", loopLength)
	fmt.Println("Pre-loop cycles:", loopStart)
	remainingCycles := (cycleCount - loopStart) % loopLength
	fmt.Println("Remaining post-loop cycles:", remainingCycles)
	for i := 0; i < remainingCycles; i++ {
		grid.SpinCycle()
	}

	//fmt.Println("After spin:")
	//fmt.Println(grid)
	fmt.Println("Total load afrer spin:", grid.TotalLoad())
}
