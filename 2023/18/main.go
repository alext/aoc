package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

type Pos = helpers.Pos

type Direction uint8

const (
	Right Direction = iota
	Down
	Left
	Up
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
		return Pos{Y: 1}
	case Down:
		return Pos{Y: -1}
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

func ParseInstruction(input string) (Instruction, Instruction) {
	i1 := Instruction{}
	i2 := Instruction{}
	var dirStr string

	// R 6 (#70c710)
	_, err := fmt.Sscanf(input, "%s %d (#%5x%d)", &dirStr, &i1.Length, &i2.Length, &i2.Dir)
	if err != nil {
		log.Fatalln("Failed to parse input", input, "err:", err)
	}
	i1.Dir = ParseDirection(dirStr)

	return i1, i2
}

func (i Instruction) Delta() Pos {
	switch i.Dir {
	case Up:
		return Pos{Y: i.Length}
	case Down:
		return Pos{Y: -i.Length}
	case Left:
		return Pos{X: -i.Length}
	case Right:
		return Pos{X: i.Length}
	default:
		log.Fatalln("Unexpected direction", i.Dir)
		return Pos{}
	}
}

type Lagoon struct {
	Corners []Pos
	Digger  Pos
}

func NewLagoon() *Lagoon {
	return &Lagoon{
		Corners: []Pos{{X: 0, Y: 0}},
	}
}

func (l *Lagoon) FillCorners(digPlan []Instruction) {
	for i, inst := range digPlan {
		l.Digger = l.Digger.Add(inst.Delta())
		if i == len(digPlan)-1 {
			if l.Digger != (Pos{}) {
				log.Fatalln("Final position not 0,0 - was", l.Digger)
			}
		} else {
			l.Corners = append(l.Corners, l.Digger)
		}
	}
}

func (l *Lagoon) TotalArea() int {
	// Shoelace + Pick's theorum
	sum := 0
	boundary := 0
	p0 := l.Corners[len(l.Corners)-1]
	for _, p1 := range l.Corners {
		sum += p0.Y*p1.X - p0.X*p1.Y
		boundary += p0.DistanceTo(p1)
		p0 = p1
	}
	//fmt.Printf("Sum: %d, boundary: %d\n", sum, boundary)
	return sum/2 + boundary/2 + 1
}

func main() {
	var instructions, instructions2 []Instruction
	helpers.ScanLines(os.Stdin, func(line string) {
		inst1, inst2 := ParseInstruction(line)
		instructions = append(instructions, inst1)
		instructions2 = append(instructions2, inst2)
	})

	fmt.Println(instructions)
	fmt.Println(instructions2)

	lagoon := NewLagoon()
	lagoon.FillCorners(instructions)
	//fmt.Println("Corners:", lagoon.Corners)
	fmt.Println("Area:", lagoon.TotalArea())

	// Part 2
	lagoon = NewLagoon()
	lagoon.FillCorners(instructions2)
	//fmt.Println("Corners:", lagoon.Corners)
	fmt.Println("Area 2:", lagoon.TotalArea())
}
