package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Pos = helpers.Pos

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

func ParseDirection(d string) Direction {
	switch d {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	default:
		panic("Unrecognised direction " + d)
	}
}

func (d Direction) Delta() Pos {
	switch d {
	case Up:
		return Pos{Y: -1}
	case Down:
		return Pos{Y: 1}
	case Left:
		return Pos{X: -1}
	case Right:
		return Pos{X: 1}
	default:
		log.Fatalln("Unexpected direction", d)
		return Pos{}
	}
}

type Instruction struct {
	Dir    Direction
	Length int
}

var lineRe = regexp.MustCompile(`^([UDRL])\s+(\d+)\s+`)

func ParseInstruction(input string) Instruction {
	matches := lineRe.FindStringSubmatch(input)
	if matches == nil {
		log.Fatalln("Failed to parse instruction", input)
	}
	return Instruction{
		Dir:    ParseDirection(matches[1]),
		Length: helpers.MustAtoi(matches[2]),
	}
}

type Lagoon struct {
	Squares   map[Pos]bool
	MinExtent Pos
	MaxExtent Pos
	Digger    Pos
}

func (l Lagoon) String() string {
	var b strings.Builder
	for y := l.MinExtent.Y; y <= l.MaxExtent.Y; y++ {
		for x := l.MinExtent.X; x <= l.MaxExtent.X; x++ {
			if x == 0 && y == 0 {
				b.WriteString("0")
			} else if l.Squares[Pos{X: x, Y: y}] {
				b.WriteString("#")
			} else {
				b.WriteString(".")
			}
		}
		if y < l.MaxExtent.Y {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func NewLagoon() *Lagoon {
	return &Lagoon{
		Squares: map[Pos]bool{{X: 0, Y: 0}: true},
	}
}

func (l *Lagoon) DigSquare(pos Pos) {
	l.Squares[pos] = true
	// Add 1 offset to ensure there's an empty border all round
	l.MinExtent.X = min(l.MinExtent.X, pos.X-1)
	l.MinExtent.Y = min(l.MinExtent.Y, pos.Y-1)
	l.MaxExtent.X = max(l.MaxExtent.X, pos.X+1)
	l.MaxExtent.Y = max(l.MaxExtent.Y, pos.Y+1)
}

func (l *Lagoon) DigTrenches(digPlan []Instruction) {
	for _, inst := range digPlan {
		delta := inst.Dir.Delta()
		for i := 0; i < inst.Length; i++ {
			l.Digger = l.Digger.Add(delta)
			l.DigSquare(l.Digger)
		}
	}
}

func (l *Lagoon) FillLoop() {
	outsideSquares := map[Pos]bool{l.MinExtent: true}
	candidates := []Pos{l.MinExtent}
	for len(candidates) > 0 {
		sq := candidates[0]
		candidates = candidates[1:]

		for _, neighbour := range []Pos{
			{X: sq.X, Y: sq.Y + 1},
			{X: sq.X, Y: sq.Y - 1},
			{X: sq.X + 1, Y: sq.Y},
			{X: sq.X - 1, Y: sq.Y},
		} {
			if neighbour.X < l.MinExtent.X || neighbour.X > l.MaxExtent.X ||
				neighbour.Y < l.MinExtent.Y || neighbour.Y > l.MaxExtent.Y {
				continue
			}
			if l.Squares[neighbour] {
				continue
			}
			if outsideSquares[neighbour] {
				// Previously looked at
				continue
			}
			outsideSquares[neighbour] = true
			candidates = append(candidates, neighbour)
		}
	}

	for y := l.MinExtent.Y; y <= l.MaxExtent.Y; y++ {
		for x := l.MinExtent.X; x <= l.MaxExtent.X; x++ {
			pos := Pos{X: x, Y: y}
			if l.Squares[pos] || outsideSquares[pos] {
				continue
			}
			l.DigSquare(pos)
		}
	}
}

func main() {
	var instructions []Instruction
	helpers.ScanLines(os.Stdin, func(line string) {
		instructions = append(instructions, ParseInstruction(line))
	})

	lagoon := NewLagoon()
	lagoon.DigTrenches(instructions)

	fmt.Println(lagoon)

	fmt.Println("Dug squares:", len(lagoon.Squares))

	lagoon.FillLoop()

	fmt.Println(lagoon)

	fmt.Println("Extent:", lagoon.MinExtent, lagoon.MaxExtent)
	fmt.Println("Dug squares:", len(lagoon.Squares))
}
