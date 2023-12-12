package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Pipe struct {
	Row   int
	Col   int
	North bool
	East  bool
	South bool
	West  bool
	Start bool

	Visited  bool
	Distance int
}

func ParsePipe(ch string, row, col int) *Pipe {
	p := &Pipe{Row: row, Col: col}
	switch ch {
	case `|`:
		p.North = true
		p.South = true
	case `-`:
		p.East = true
		p.West = true
	case `L`:
		p.North = true
		p.East = true
	case `J`:
		p.North = true
		p.West = true
	case `7`:
		p.South = true
		p.West = true
	case `F`:
		p.South = true
		p.East = true
	case `S`:
		p.Start = true
	case `.`:
		return nil
	default:
		panic("Unexpected character: " + ch)
	}
	return p
}

type Grid struct {
	Pipes [][]*Pipe
	Start *Pipe
}

func (g Grid) PipeAt(row, col int) *Pipe {
	if row < 0 || row >= len(g.Pipes) {
		return nil
	}
	if col < 0 || col >= len(g.Pipes[row]) {
		return nil
	}
	return g.Pipes[row][col]
}

func (g Grid) getStartNeighbours() []*Pipe {
	var neighbours []*Pipe
	if p := g.PipeAt(g.Start.Row-1, g.Start.Col); p != nil && p.South {
		neighbours = append(neighbours, p)
	}
	if p := g.PipeAt(g.Start.Row+1, g.Start.Col); p != nil && p.North {
		neighbours = append(neighbours, p)
	}
	if p := g.PipeAt(g.Start.Row, g.Start.Col-1); p != nil && p.East {
		neighbours = append(neighbours, p)
	}
	if p := g.PipeAt(g.Start.Row, g.Start.Col+1); p != nil && p.West {
		neighbours = append(neighbours, p)
	}

	if len(neighbours) != 2 {
		log.Fatalf("Found %d start neighbours", len(neighbours))
	}
	return neighbours
}

func (g Grid) getNext(p *Pipe) *Pipe {
	if p.North {
		next := g.PipeAt(p.Row-1, p.Col)
		if next != nil && !next.Visited {
			return next
		}
	}
	if p.South {
		next := g.PipeAt(p.Row+1, p.Col)
		if next != nil && !next.Visited {
			return next
		}
	}
	if p.East {
		next := g.PipeAt(p.Row, p.Col+1)
		if next != nil && !next.Visited {
			return next
		}
	}
	if p.West {
		next := g.PipeAt(p.Row, p.Col-1)
		if next != nil && !next.Visited {
			return next
		}
	}
	return nil
}

func (g Grid) MaxDistance() int {
	g.Start.Visited = true

	current := g.getStartNeighbours()
	for _, p := range current {
		p.Visited = true
		p.Distance = 1
	}

	var next []*Pipe
	maxDistance := 0
	for len(current) > 0 {
		for _, p := range current {
			if p.Distance > maxDistance {
				maxDistance = p.Distance
			}
			n := g.getNext(p)
			if n == nil {
				continue
			}
			n.Visited = true
			n.Distance = p.Distance + 1
			next = append(next, n)
		}
		current = next
		next = nil
	}

	return maxDistance
}

func main() {

	var g Grid
	var rowNum int
	helpers.ScanLines(os.Stdin, func(line string) {
		row := make([]*Pipe, 0, len(line))
		for colNum, ch := range strings.Split(line, "") {
			p := ParsePipe(ch, rowNum, colNum)
			if p != nil && p.Start {
				g.Start = p
			}
			row = append(row, p)
		}
		g.Pipes = append(g.Pipes, row)
		rowNum++
	})

	fmt.Println("Max distance:", g.MaxDistance())
}
