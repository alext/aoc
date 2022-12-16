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

func ParsePos(in string) Pos {
	x, y, _ := strings.Cut(in, ",")
	return Pos{
		X: helpers.MustAtoi(x),
		Y: helpers.MustAtoi(y),
	}
}

type Grid struct {
	Positions map[Pos]string
	MinX      int
	MaxX      int
	MaxY      int
	SandCount int
}

func NewGrid() *Grid {
	return &Grid{
		Positions: make(map[Pos]string),
		MinX:      500,
		MaxX:      500,
	}
}

func (g *Grid) String() string {
	b := &strings.Builder{}
	for y := 0; y <= g.MaxY; y++ {
		for x := g.MinX; x <= g.MaxX; x++ {
			if ch, ok := g.Positions[Pos{X: x, Y: y}]; ok {
				b.WriteString(ch)
			} else if x == 500 && y == 0 {
				b.WriteString("+")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (g *Grid) setPos(p Pos, v string) {
	g.Positions[p] = v
	if p.X < g.MinX {
		g.MinX = p.X
	}
	if p.X > g.MaxX {
		g.MaxX = p.X
	}
	if p.Y > g.MaxY {
		g.MaxY = p.Y
	}
	if v == "o" {
		g.SandCount++
	}
}

func (g *Grid) AddRock(corners []Pos) {
	previous := corners[0]
	corners = corners[1:]
	g.setPos(previous, "#")

	for _, corner := range corners {
		if corner.X == previous.X {
			for y := helpers.Min(previous.Y, corner.Y); y <= helpers.Max(previous.Y, corner.Y); y++ {
				g.setPos(Pos{X: corner.X, Y: y}, "#")
			}
		} else if corner.Y == previous.Y {
			for x := helpers.Min(previous.X, corner.X); x <= helpers.Max(previous.X, corner.X); x++ {
				g.setPos(Pos{X: x, Y: corner.Y}, "#")
			}
		} else {
			log.Fatalln("corners are diagonal", previous, corner)
		}
		previous = corner
	}
}

func (g *Grid) AddSand() bool {
	sandPos := Pos{X: 500, Y: 0}
	for sandPos.Y <= g.MaxY {
		testPos := Pos{X: sandPos.X, Y: sandPos.Y + 1}
		if _, found := g.Positions[testPos]; !found {
			sandPos = testPos
			continue
		}
		testPos.X = sandPos.X - 1
		if _, found := g.Positions[testPos]; !found {
			sandPos = testPos
			continue
		}
		testPos.X = sandPos.X + 1
		if _, found := g.Positions[testPos]; !found {
			sandPos = testPos
			continue
		}
		g.setPos(sandPos, "o")
		return true
	}
	return false
}

func main() {

	g := NewGrid()

	helpers.ScanLines(os.Stdin, func(line string) {
		var corners []Pos
		for _, corner := range strings.Split(line, " -> ") {
			corners = append(corners, ParsePos(corner))
		}
		g.AddRock(corners)
	})

	fmt.Println(g)
	for {
		landed := g.AddSand()
		if !landed {
			break
		}
	}
	fmt.Println(g)
	fmt.Println("SandCount", g.SandCount)
}
