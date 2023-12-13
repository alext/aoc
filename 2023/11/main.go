package main

import (
	"fmt"
	"os"
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

func expandRow(row []string, emptyCols map[int]bool) []string {
	var expanded []string
	for col, ch := range row {
		expanded = append(expanded, ch)
		if emptyCols[col] {
			expanded = append(expanded, ch)
		}
	}
	return expanded
}

func expandGalaxy(input RawGalaxy) RawGalaxy {
	if len(input) == 0 {
		return [][]string{}
	}
	emptyCols := make(map[int]bool)
colLoop:
	for col := 0; col < len(input[0]); col++ {
		for row := 0; row < len(input); row++ {
			if input[row][col] == "#" {
				continue colLoop
			}
		}
		emptyCols[col] = true
	}
	var expanded RawGalaxy
rowLoop:
	for row := 0; row < len(input); row++ {
		expanddedRow := expandRow(input[row], emptyCols)
		expanded = append(expanded, expanddedRow)
		for _, ch := range input[row] {
			if ch == "#" {
				continue rowLoop
			}
		}
		// No galaxies in row, so append again
		expanded = append(expanded, expanddedRow)
	}
	return expanded
}

func main() {
	var rawGalaxy RawGalaxy
	helpers.ScanLines(os.Stdin, func(line string) {
		rawGalaxy = append(rawGalaxy, strings.Split(line, ""))
	})
	fmt.Println(rawGalaxy)
	rawGalaxy = expandGalaxy(rawGalaxy)
	fmt.Println(rawGalaxy)

	var galaxies []helpers.Pos
	for row := range rawGalaxy {
		for col, ch := range rawGalaxy[row] {
			if ch == "#" {
				galaxies = append(galaxies, helpers.Pos{X: col, Y: row})
			}
		}
	}
	fmt.Println(galaxies)
	totalDistance := 0
	for i, a := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]
			totalDistance += a.DistanceTo(b)
		}
	}
	fmt.Println("Total distance:", totalDistance)
}
