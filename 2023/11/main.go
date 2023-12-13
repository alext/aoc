package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/alext/aoc/helpers"
)

type RawGalaxy [][]string

func (g RawGalaxy) String() string {
	b := strings.Builder{}
	for _, row := range g {
		b.WriteString(strings.Join(row, ""))
		b.WriteString("\n")
	}
	return b.String()
}

func (g RawGalaxy) EmptyRowsCols() ([]int, []int) {
	if len(g) == 0 {
		return nil, nil
	}
	var emptyRows []int
rowLoop:
	for row := 0; row < len(g); row++ {
		for _, ch := range g[row] {
			if ch == "#" {
				continue rowLoop
			}
		}
		emptyRows = append(emptyRows, row)
	}
	var emptyCols []int
colLoop:
	for col := 0; col < len(g[0]); col++ {
		for row := 0; row < len(g); row++ {
			if g[row][col] == "#" {
				continue colLoop
			}
		}
		emptyCols = append(emptyCols, col)
	}
	return emptyRows, emptyCols
}

type Galaxies []helpers.Pos

func (g Galaxies) MakeExpanded(emptyRows, emptyCols []int, factor int) Galaxies {
	expanded := slices.Clone(g)

	for r := len(emptyRows) - 1; r >= 0; r-- {
		for i := range expanded {
			if expanded[i].Y > emptyRows[r] {
				expanded[i].Y += factor - 1
			}
		}
	}
	for c := len(emptyCols) - 1; c >= 0; c-- {
		for i := range expanded {
			if expanded[i].X > emptyCols[c] {
				expanded[i].X += factor - 1
			}
		}
	}

	return expanded
}

func (g Galaxies) TotalDistance() int {
	totalDistance := 0
	for i, a := range g {
		for j := i + 1; j < len(g); j++ {
			b := g[j]
			totalDistance += a.DistanceTo(b)
		}
	}
	return totalDistance
}

func main() {
	var rawGalaxy RawGalaxy
	helpers.ScanLines(os.Stdin, func(line string) {
		rawGalaxy = append(rawGalaxy, strings.Split(line, ""))
	})
	fmt.Println(rawGalaxy)

	var galaxies Galaxies
	for row := range rawGalaxy {
		for col, ch := range rawGalaxy[row] {
			if ch == "#" {
				galaxies = append(galaxies, helpers.Pos{X: col, Y: row})
			}
		}
	}

	emptyRows, emptyCols := rawGalaxy.EmptyRowsCols()
	expanded := galaxies.MakeExpanded(emptyRows, emptyCols, 2)

	fmt.Println(galaxies)
	fmt.Println(expanded)

	fmt.Println("Total distance:", expanded.TotalDistance())

	expanded = galaxies.MakeExpanded(emptyRows, emptyCols, 1_000_000)
	fmt.Println("Total distance 2:", expanded.TotalDistance())
}
