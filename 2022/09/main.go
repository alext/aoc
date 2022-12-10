package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Pos struct {
	X int
	Y int
}

type Rope struct {
	Head        Pos
	Tail        Pos
	TailVisited map[Pos]bool
}

func NewRope() *Rope {
	r := &Rope{
		TailVisited: make(map[Pos]bool),
	}
	r.TailVisited[r.Tail] = true
	return r
}

func (r Rope) String() string {
	return fmt.Sprintf("Head (%d,%d) Tail (%d,%d)", r.Head.X, r.Head.Y, r.Tail.X, r.Tail.Y)
}
func absSign(n int) (int, int) {
	if n >= 0 {
		return n, 1
	}
	return -n, -1
}

func (r *Rope) moveTail() {
	xDelta, xSign := absSign(r.Head.X - r.Tail.X)
	yDelta, ySign := absSign(r.Head.Y - r.Tail.Y)

	if xDelta < 2 && yDelta < 2 {
		// No move needed
		return
	}

	if xDelta >= 2 {
		r.Tail.X += xSign
		if yDelta > 0 {
			r.Tail.Y += ySign
		}
	}
	if yDelta >= 2 {
		r.Tail.Y += ySign
		if xDelta > 0 {
			r.Tail.X += xSign
		}
	}
	r.TailVisited[r.Tail] = true
}

func (r *Rope) moveOne(direction string) {
	switch direction {
	case "U":
		r.Head.Y++
	case "D":
		r.Head.Y--
	case "R":
		r.Head.X++
	case "L":
		r.Head.X--
	default:
		log.Fatalln("Unexpected direction:", direction)
	}
	r.moveTail()
}

func (r *Rope) Move(direction string, amount int) {
	for i := 0; i < amount; i++ {
		r.moveOne(direction)
	}
}

func main() {
	r := NewRope()
	fmt.Println("Starting:", r)

	helpers.ScanLines(os.Stdin, func(line string) {
		dir, amount, _ := strings.Cut(line, " ")
		r.Move(dir, helpers.MustAtoi(amount))
	})

	fmt.Println("Visited squares:", len(r.TailVisited))
}
