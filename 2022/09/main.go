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

func (p Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Rope struct {
	Head        Pos
	Tails       []Pos
	TailVisited map[Pos]bool
}

func NewRope(numTails int) *Rope {
	r := &Rope{
		Tails:       make([]Pos, numTails),
		TailVisited: make(map[Pos]bool),
	}
	r.TailVisited[Pos{0, 0}] = true
	return r
}

func (r Rope) String() string {
	return fmt.Sprintf("Head %s Tail %s", r.Head, r.Tails)
}
func absSign(n int) (int, int) {
	if n >= 0 {
		return n, 1
	}
	return -n, -1
}

func (r *Rope) moveTail(index int) bool {
	var current, previous *Pos
	current = &r.Tails[index]
	if index == 0 {
		previous = &r.Head
	} else {
		previous = &r.Tails[index-1]
	}

	xDelta, xSign := absSign(previous.X - current.X)
	yDelta, ySign := absSign(previous.Y - current.Y)
	if xDelta < 2 && yDelta < 2 {
		// No move needed
		return false
	}

	if xDelta >= 2 {
		current.X += xSign
		if yDelta > 0 {
			current.Y += ySign
		}
	} else if yDelta >= 2 {
		current.Y += ySign
		if xDelta > 0 {
			current.X += xSign
		}
	}
	return true
}

func (r *Rope) moveTails() {
	for i := range r.Tails {
		moved := r.moveTail(i)
		if !moved {
			return
		}
	}
	r.TailVisited[r.Tails[len(r.Tails)-1]] = true
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
	r.moveTails()
}

func (r *Rope) Move(direction string, amount int) {
	for i := 0; i < amount; i++ {
		r.moveOne(direction)
	}
}

func main() {
	r := NewRope(1)
	r2 := NewRope(9)

	helpers.ScanLines(os.Stdin, func(line string) {
		dir, amount, _ := strings.Cut(line, " ")
		r.Move(dir, helpers.MustAtoi(amount))
		r2.Move(dir, helpers.MustAtoi(amount))
	})

	fmt.Println("Visited squares r :", len(r.TailVisited))
	fmt.Println("Visited squares r2:", len(r2.TailVisited))

}
