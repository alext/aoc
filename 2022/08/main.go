package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Grid [][]int

func (g *Grid) AddRow(rowStr string) {
	row := make([]int, 0)
	for _, height := range strings.Split(rowStr, "") {
		row = append(row, helpers.MustAtoi(height))
	}
	if len(*g) > 0 && len((*g)[0]) != len(row) {
		log.Fatal("Row length mismatch")
	}
	*g = append(*g, row)
}

func (g Grid) TreeVisible(row, col int) bool {
	t := g[row][col]
	visibleDirections := 4
	for r := 0; r < row; r++ {
		if g[r][col] >= t {
			visibleDirections--
			break
		}
	}
	for r := row + 1; r < len(g); r++ {
		if g[r][col] >= t {
			visibleDirections--
			break
		}
	}
	for c := 0; c < col; c++ {
		if g[row][c] >= t {
			visibleDirections--
			break
		}
	}
	for c := col + 1; c < len(g[row]); c++ {
		if g[row][c] >= t {
			visibleDirections--
			break
		}
	}
	return visibleDirections > 0
}

func (g Grid) ScenicScore(row, col int) int {
	t := g[row][col]
	north := 0
	for r := row - 1; r >= 0; r-- {
		north++
		if g[r][col] >= t {
			break
		}
	}
	south := 0
	for r := row + 1; r < len(g); r++ {
		south++
		if g[r][col] >= t {
			break
		}
	}
	east := 0
	for c := col - 1; c >= 0; c-- {
		east++
		if g[row][c] >= t {
			break
		}
	}
	west := 0
	for c := col + 1; c < len(g[row]); c++ {
		west++
		if g[row][c] >= t {
			break
		}
	}
	return north * south * east * west
}

func (g Grid) VisibleCount() int {
	count := 0
	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			if g.TreeVisible(row, col) {
				count++
			}
		}
	}
	return count
}

func (g Grid) BestScenicScore() int {
	best := 0
	for row := 0; row < len(g); row++ {
		for col := 0; col < len(g[row]); col++ {
			score := g.ScenicScore(row, col)
			if score > best {
				best = score
			}
		}
	}
	return best
}

func main() {
	g := make(Grid, 0)

	helpers.ScanLines(os.Stdin, func(line string) { g.AddRow(line) })

	fmt.Println("Visible count:", g.VisibleCount())

	fmt.Println("Best scenic score:", g.BestScenicScore())
}
