package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Shape [][]string

func (s Shape) String() string {
	var b = strings.Builder{}
	for i, line := range s {
		b.WriteString(strings.Join(line, ""))
		if i < len(s)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func (s Shape) isReflectionCol(col int) bool {
	if col >= len(s[0])-1 {
		return false
	}
	leftCol, rightCol := col, col+1
	for leftCol >= 0 && rightCol < len(s[0]) {
		for row := 0; row < len(s); row++ {
			if s[row][leftCol] != s[row][rightCol] {
				return false
			}
		}
		leftCol--
		rightCol++
	}
	return true
}

func (s Shape) isReflectionRow(row int) bool {
	if row >= len(s)-1 {
		return false
	}
	topRow, bottomRow := row, row+1
	for topRow >= 0 && bottomRow < len(s) {
		for col := 0; col < len(s[topRow]); col++ {
			if s[topRow][col] != s[bottomRow][col] {
				return false
			}
		}
		topRow--
		bottomRow++
	}
	return true
}

func (s Shape) SummariseReflections() int {
	if len(s) == 0 {
		return 0
	}
	summary := 0
	for col := 0; col < len(s[0]); col++ {
		if s.isReflectionCol(col) {
			fmt.Println("Found col reflection:", col)
			summary += col + 1
		}
	}
	for row := 0; row < len(s); row++ {
		if s.isReflectionRow(row) {
			summary += (row + 1) * 100
		}
	}
	return summary
}

func main() {

	var shapes []Shape
	var s Shape
	helpers.ScanLines(os.Stdin, func(line string) {
		if line == "" {
			shapes = append(shapes, s)
			s = nil
			return
		}
		s = append(s, strings.Split(line, ""))
	})
	if s != nil {
		shapes = append(shapes, s)
	}

	totalSummary := 0
	for _, shape := range shapes {
		fmt.Println(shape)
		sum := shape.SummariseReflections()
		fmt.Println("Summary:", sum)
		totalSummary += sum
	}
	fmt.Println("Total summary:", totalSummary)
}
