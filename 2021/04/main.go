package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Board [5][5]struct {
	Number int
	Marked bool
}

func BuildBoard(input [][]int) *Board {
	if len(input) != 5 || len(input[0]) != 5 {
		log.Fatal("Expected 5x5 input, got", input)
	}
	var b Board

	for x, row := range input {
		for y, n := range row {
			b[x][y].Number = n
		}
	}
	return &b
}

func (b *Board) String() string {
	var s strings.Builder
	for _, row := range b {
		for _, square := range row {
			fmt.Fprintf(&s, "%2d", square.Number)
			if square.Marked {
				s.WriteString("* ")
			} else {
				s.WriteString("  ")
			}
		}
		s.WriteString("\n")
	}
	return s.String()
}

func (b *Board) HasWon() bool {
	// Check rows
	for x := 0; x < 5; x++ {
		won := true
		for y := 0; y < 5; y++ {
			if !b[x][y].Marked {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}
	// Check columns
	for y := 0; y < 5; y++ {
		won := true
		for x := 0; x < 5; x++ {
			if !b[x][y].Marked {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}
	return false
}

func (b *Board) Play(call int) bool {
	// TODO: do we need to allow for a number appearing more than once on a board?
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b[x][y].Number == call {
				b[x][y].Marked = true
				return b.HasWon()
			}
		}
	}
	return false
}

func (b *Board) Score(lastCall int) int {
	score := 0
	for _, row := range b {
		for _, square := range row {
			if !square.Marked {
				score += square.Number
			}
		}
	}
	return score * lastCall
}

func main() {
	var calls []int

	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data)-1; i++ {
			if string(data[i:i+2]) == "\n\n" {
				return i + 2, data[:i], nil
			}
		}
		if atEOF {
			return len(data), data, bufio.ErrFinalToken
		}
		// request more data
		return 0, nil, nil
	}

	var boards []*Board
	helpers.ScanWrapper(os.Stdin, splitFunc, func(section string) {
		var numbers [][]int
		if calls == nil {
			// first section with all the calls
			for _, num := range helpers.SplitCSV(section) {
				calls = append(calls, helpers.MustAtoi(num))
			}
			return
		}

		helpers.ScanLines(strings.NewReader(section), func(line string) {
			var lineNums []int

			helpers.ScanWrapper(strings.NewReader(line), bufio.ScanWords, func(num string) {
				lineNums = append(lineNums, helpers.MustAtoi(num))
			})
			numbers = append(numbers, lineNums)
		})
		boards = append(boards, BuildBoard(numbers))
	})

	winningScore := 0
Loop:
	for _, call := range calls {
		for _, b := range boards {
			if b.Play(call) {
				fmt.Println("Winning board")
				fmt.Println(b)
				winningScore = b.Score(call)
				break Loop
			}
		}
	}
	if winningScore != 0 {
		fmt.Println("Winning score:", winningScore)
	}
}
