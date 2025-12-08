package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/alext/aoc/helpers"
)

type Dial struct {
	Pos      int
	AtZero   int
	PassZero int
}

func (d Dial) String() string {
	return fmt.Sprintf("%2d [%d,%d]", d.Pos, d.AtZero, d.PassZero)
}

func (d *Dial) Move(move int) {
	startedAtZero := d.Pos == 0
	d.Pos += move
	for d.Pos >= 100 {
		d.Pos -= 100
		if d.Pos != 0 {
			d.PassZero++
		}
	}
	for d.Pos < 0 {
		d.Pos += 100
		if d.Pos < 0 || !startedAtZero {
			d.PassZero++
		}
	}
	if d.Pos == 0 {
		d.AtZero++
		d.PassZero++
	}
}

func main() {
	lineRe := regexp.MustCompile(`([LR])(\d+)`)
	var moves []int
	helpers.ScanLines(os.Stdin, func(line string) {
		parts := lineRe.FindStringSubmatch(line)

		move := helpers.MustAtoi(parts[2])
		if parts[1] == "L" {
			move = -move
		}
		moves = append(moves, move)
	})

	// fmt.Println(moves)

	d := Dial{Pos: 50}

	for _, move := range moves {
		d.Move(move)
		// fmt.Println("Move", move, d)
	}

	fmt.Println("Number of times at zero:", d.AtZero)
	fmt.Println("Number of times at or passing zero:", d.PassZero)
}
