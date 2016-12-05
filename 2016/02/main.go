package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

var numberpad [5][5]string

func init() {
	numberpad = [5][5]string{
		{" ", " ", "D", " ", " "},
		{" ", "A", "B", "C", " "},
		{"5", "6", "7", "8", "9"},
		{" ", "2", "3", "4", " "},
		{" ", " ", "1", " ", " "},
	}
}

type Position struct {
	X int
	Y int
}

func (p Position) Move(direction rune) Position {
	newpos := p
	switch direction {
	case 'U':
		if newpos.Y < len(numberpad[0])-1 {
			newpos.Y += 1
		}
	case 'D':
		if newpos.Y > 0 {
			newpos.Y -= 1
		}
	case 'L':
		if newpos.X > 0 {
			newpos.X -= 1
		}
	case 'R':
		if newpos.X < len(numberpad[0])-1 {
			newpos.X += 1
		}
	}
	if numberpad[newpos.Y][newpos.X] == " " {
		return p
	}
	return newpos
}

var StartingPos = Position{X: 1, Y: 2}

func main() {
	helpers.ScanLines(os.Stdin, func(line string) {
		pos := StartingPos
		for _, dir := range line {
			pos = pos.Move(dir)
			//fmt.Printf("X: %d, Y: %d\n", pos.X, pos.Y)
		}
		fmt.Println("Number:", numberpad[pos.Y][pos.X])
	})
}
