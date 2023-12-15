package main

import (
	"fmt"
	"log"
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

func (s Shape) isReflectionCol(col int) (bool, bool) {
	if col >= len(s[0])-1 {
		return false, false
	}
	smudgeFound := false
	leftCol, rightCol := col, col+1
	for leftCol >= 0 && rightCol < len(s[0]) {
		for row := 0; row < len(s); row++ {
			if s[row][leftCol] != s[row][rightCol] {
				if smudgeFound {
					return false, false
				} else {
					smudgeFound = true
				}
			}
		}
		leftCol--
		rightCol++
	}
	if smudgeFound {
		return false, true
	}
	return true, false
}

func (s Shape) isReflectionRow(row int) (bool, bool) {
	if row >= len(s)-1 {
		return false, false
	}
	smudgeFound := false
	topRow, bottomRow := row, row+1
	for topRow >= 0 && bottomRow < len(s) {
		for col := 0; col < len(s[topRow]); col++ {
			if s[topRow][col] != s[bottomRow][col] {
				if smudgeFound {
					return false, false
				} else {
					smudgeFound = true
				}
			}
		}
		topRow--
		bottomRow++
	}
	if smudgeFound {
		return false, true
	}
	return true, false
}

func (s Shape) SummariseReflections() int {
	if len(s) == 0 {
		return 0
	}
	summary := 0
	for col := 0; col < len(s[0]); col++ {
		if is, _ := s.isReflectionCol(col); is {
			//fmt.Println("Found reflection col:", col)
			summary += col + 1
		}
	}
	for row := 0; row < len(s); row++ {
		if is, _ := s.isReflectionRow(row); is {
			//fmt.Println("Found reflection row:", row)
			summary += (row + 1) * 100
		}
	}
	return summary
}

func (s Shape) SummariseSmudgeReflections() int {
	for col := 0; col < len(s[0]); col++ {
		found, smudgeFound := s.isReflectionCol(col)
		if found {
			continue
		}
		if smudgeFound {
			//fmt.Printf("Found smudge from col %d at %d, %d\n", col, smudge.X, smudge.Y)
			return col + 1
		}
	}
	for row := 0; row < len(s); row++ {
		found, smudgeFound := s.isReflectionRow(row)
		if found {
			continue
		}
		if smudgeFound {
			//fmt.Printf("Found smudge from row %d at %d, %d\n", row, smudge.X, smudge.Y)
			return (row + 1) * 100
		}
	}
	log.Fatalln("Failed to find smudge in\n", s)
	return -1
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
		//fmt.Println(shape)
		sum := shape.SummariseReflections()
		//fmt.Println("Summary:", sum)
		totalSummary += sum
	}
	fmt.Println("Total summary:", totalSummary)

	totalSummary = 0
	for _, shape := range shapes {
		//fmt.Println(shape)
		sum := shape.SummariseSmudgeReflections()
		//fmt.Println("Summary:", sum)
		totalSummary += sum
	}
	fmt.Println("Total smudge line summary:", totalSummary)
}
