package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Pipe struct {
	ch    string
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
	p := &Pipe{ch: ch, Row: row, Col: col}
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

func (g Grid) String() string {
	var s strings.Builder
	for _, row := range g.Pipes {
		for _, p := range row {
			if p == nil {
				s.WriteString(`.`)
			} else if p.Start {
				s.WriteString("S")
			} else {
				s.WriteString(p.ch)
			}
		}
		s.WriteString("\n")
	}
	return s.String()
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
		g.Start.North = true
	}
	if p := g.PipeAt(g.Start.Row+1, g.Start.Col); p != nil && p.North {
		neighbours = append(neighbours, p)
		g.Start.South = true
	}
	if p := g.PipeAt(g.Start.Row, g.Start.Col-1); p != nil && p.East {
		neighbours = append(neighbours, p)
		g.Start.West = true
	}
	if p := g.PipeAt(g.Start.Row, g.Start.Col+1); p != nil && p.West {
		neighbours = append(neighbours, p)
		g.Start.East = true
	}
	switch {
	case g.Start.North && g.Start.South:
		g.Start.ch = "|"
	case g.Start.East && g.Start.West:
		g.Start.ch = "-"
	case g.Start.North && g.Start.East:
		g.Start.ch = "L"
	case g.Start.North && g.Start.West:
		g.Start.ch = "J"
	case g.Start.South && g.Start.East:
		g.Start.ch = "F"
	case g.Start.South && g.Start.West:
		g.Start.ch = "7"
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

func (g Grid) ClearNonLoopPipe() {
	for r, row := range g.Pipes {
		for c, p := range row {
			if p == nil {
				continue
			}
			if p.Start || p.Visited {
				continue
			}
			g.Pipes[r][c] = nil
		}
	}
}

func (g Grid) IsInsideLoop(row, col int) bool {
	if g.Pipes[row][col] != nil {
		// On the loop - therefore not inside
		return false
	}
	gridRow := g.Pipes[row]
	crossings := 0
	for i := col - 1; i >= 0; i-- {
		p := gridRow[i]
		if p == nil {
			continue
		}
		segmentStart := ""
		switch p.ch {
		case "|":
			crossings++
		case "J", "7":
			segmentStart = p.ch
		default:
			log.Fatalln("Open-ended pipe:", p.ch)
		}
		for segmentStart != "" {
			// We're dealing with a horizontal pipe segment. Loop for the start of it.
			i--
			p = gridRow[i]
			if p == nil {
				log.Fatalf("Dangling segment at r:%d, c:%d from row:%d, col:%d", row, i, row, col)
			}
			switch p.ch {
			case "-":
				continue
			case "F":
				if segmentStart == "J" {
					crossings++
				}
				segmentStart = ""
			case "L":
				if segmentStart == "7" {
					crossings++
				}
				segmentStart = ""
			default:
				log.Fatalln("Mismatched segment end:", p.ch)
			}
		}
	}
	return crossings%2 == 1
}

func (g Grid) CountInsideSquares() int {
	count := 0
	for r, row := range g.Pipes {
		for c := range row {
			if g.IsInsideLoop(r, c) {
				count++
			}
		}
	}
	return count
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
	fmt.Println(g)
	g.ClearNonLoopPipe()
	fmt.Println(g)

	fmt.Println("Inside squares:", g.CountInsideSquares())
}
