package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/alext/aoc/helpers"
)

type Direction uint8

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) Opposite() Direction {
	return (d + 2) % 4
}

type Square struct {
	Symbol      string
	Row         int
	Col         int
	EntryPoints []Direction
}

func (s Square) String() string { return s.Symbol }

func (s Square) ExitPoints(entryPoint Direction) []Direction {
	switch s.Symbol {
	case `/`, `\`:
		var exit Direction
		// Calculate for `/`
		switch entryPoint {
		case East:
			exit = North
		case North:
			exit = East
		case West:
			exit = South
		case South:
			exit = East
		}
		// and reverse if it's `\`
		if s.Symbol == `\` {
			exit = exit.Opposite()
		}
		return []Direction{exit}
	case `-`:
		switch entryPoint {
		case North, South:
			return []Direction{East, West}
		default:
			return []Direction{entryPoint.Opposite()}
		}
	case `|`:
		switch entryPoint {
		case East, West:
			return []Direction{North, South}
		default:
			return []Direction{entryPoint.Opposite()}
		}
	default: // `.` and anything unexpected
		return []Direction{entryPoint.Opposite()}
	}
}

type Contraption [][]*Square

func BuildContraption(input [][]string) Contraption {
	c := make(Contraption, len(input))
	for row, line := range input {
		for col, symbol := range line {
			c[row] = append(c[row], &Square{
				Symbol: symbol,
				Row:    row,
				Col:    col,
			})
		}
	}
	return c
}

func (c Contraption) AdjacentSquare(s *Square, dir Direction) *Square {
	row, col := s.Row, s.Col
	switch dir {
	case North:
		row--
	case South:
		row++
	case West:
		col--
	case East:
		col++
	}
	if row < 0 || row >= len(c) {
		return nil
	}
	line := c[row]
	if col < 0 || col >= len(line) {
		return nil
	}
	return line[col]
}

func (c Contraption) ProcessBeam() {
	start := c[0][0]
	start.EntryPoints = []Direction{West}

	var current = []*Square{start}
	var next []*Square
	for len(current) > 0 {
		for _, sq := range current {
			// TODO: The last entrypoint might not be the right one to process..
			exits := sq.ExitPoints(sq.EntryPoints[len(sq.EntryPoints)-1])
			for _, dir := range exits {
				nextSq := c.AdjacentSquare(sq, dir)
				if nextSq == nil {
					// Edge of grid
					continue
				}
				if slices.Contains(nextSq.EntryPoints, dir.Opposite()) {
					continue
				}
				nextSq.EntryPoints = append(nextSq.EntryPoints, dir.Opposite())
				next = append(next, nextSq)
			}
		}
		current, next = next, nil
	}
}

func main() {
	c := BuildContraption(helpers.ScanGrid(os.Stdin, ""))

	fmt.Println(c)
}
