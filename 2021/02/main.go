package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Submarine struct {
	position int
	depth    int
	aim      int
}

func (s *Submarine) String() string {
	return fmt.Sprintf("Submarine position: %d, depth: %d, product: %d", s.position, s.depth, s.position*s.depth)
}

func (s *Submarine) Move(direction string, amount int) {
	switch direction {
	case "forward":
		s.position += amount
	case "up":
		s.depth -= amount
	case "down":
		s.depth += amount
	default:
		panic("Unexpected direction " + direction)
	}
}

func (s *Submarine) MoveWithAim(direction string, amount int) {
	switch direction {
	case "forward":
		s.position += amount
		s.depth += amount * s.aim
	case "up":
		s.aim -= amount
	case "down":
		s.aim += amount
	default:
		panic("Unexpected direction " + direction)
	}
}

func main() {

	sub := &Submarine{}
	aimSub := &Submarine{}

	helpers.ScanLines(os.Stdin, func(line string) {
		direction, amountStr, _ := strings.Cut(line, " ")

		amount := helpers.MustAtoi(amountStr)
		sub.Move(direction, amount)
		aimSub.MoveWithAim(direction, amount)
	})

	fmt.Println(sub)
	fmt.Println(aimSub)
}
