package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

func countNeighbours(grid [][]string, col, row int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ {
		if r < 0 || r >= len(grid) {
			continue
		}
		for c := col - 1; c <= col+1; c++ {
			if c < 0 || c >= len(grid[r]) {
				continue
			}
			if c == col && r == row {
				continue
			}
			if grid[r][c] != "." {
				count++
			}
		}
	}

	return count
}

func removeAvailable(grid [][]string) int {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != "@" {
				continue
			}
			neighbours := countNeighbours(grid, col, row)
			if neighbours < 4 {
				grid[row][col] = "x"
			}
		}
	}
	removed := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == "x" {
				grid[row][col] = "."
				removed++
			}
		}
	}
	return removed
}

func main() {
	grid := helpers.ScanGrid(os.Stdin, "")

	available := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != "@" {
				continue
			}
			neighbours := countNeighbours(grid, col, row)
			if neighbours < 4 {
				available++
			}
		}
	}
	fmt.Println("Available rolls:", available)

	totalRemoved := 0
	for {
		removed := removeAvailable(grid)
		if removed == 0 {
			break
		}
		totalRemoved += removed
	}
	fmt.Println("Total removed:", totalRemoved)
}
