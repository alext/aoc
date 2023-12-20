package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

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

func (s Square) IsEnergised() bool { return len(s.EntryPoints) > 0 }

func (s *Square) Reset() {
	if len(s.EntryPoints) > 0 {
		s.EntryPoints = s.EntryPoints[:0]
	}
}

func (s Square) String() string {
	if s.Symbol == `.` {
		if len(s.EntryPoints) >= 2 {
			return strconv.Itoa(len(s.EntryPoints))
		}
		if len(s.EntryPoints) == 1 {
			switch s.EntryPoints[0] {
			case North:
				return "v"
			case South:
				return "^"
			case East:
				return "<"
			case West:
				return ">"
			}
		}
	}
	return s.Symbol
}

func (s Square) ExitPoints(entryPoint Direction) []Direction {
	switch s.Symbol {
	case `/`, `\`:
		var exit Direction
		// Calculate for `/`
		switch entryPoint {
		case East:
			exit = South
		case North:
			exit = West
		case West:
			exit = North
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

func (c Contraption) String() string {
	var b = strings.Builder{}
	for i, line := range c {
		for _, square := range line {
			b.WriteString(square.String())
		}
		if i < len(c)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func (c Contraption) Reset() {
	for _, row := range c {
		for _, sq := range row {
			sq.Reset()
		}
	}
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

func (c Contraption) processBeam(start *Square, startDir Direction) {
	type Task struct {
		Square     *Square
		EntryPoint Direction
	}

	start.EntryPoints = []Direction{startDir}
	var current = []*Task{{Square: start, EntryPoint: startDir}}
	var next []*Task
	for len(current) > 0 {
		for _, task := range current {
			exits := task.Square.ExitPoints(task.EntryPoint)
			for _, dir := range exits {
				nextSq := c.AdjacentSquare(task.Square, dir)
				if nextSq == nil {
					// Edge of grid
					continue
				}
				entry := dir.Opposite()
				if slices.Contains(nextSq.EntryPoints, entry) {
					continue
				}
				nextSq.EntryPoints = append(nextSq.EntryPoints, entry)
				next = append(next, &Task{Square: nextSq, EntryPoint: entry})
			}
		}
		current, next = next, nil
	}
}

func (c Contraption) ProcessBeam() {
	c.processBeam(c[0][0], West)
}

func (c Contraption) FindBestBeam() (*Square, Direction, int) {
	var bestSq *Square
	var bestEntry Direction
	var bestCount int

	check := func(row, col int, dir Direction) {
		c.Reset()
		sq := c[row][col]
		c.processBeam(sq, dir)
		count := c.EnergisedSquares()
		//fmt.Printf("(%d,%d) from %d, count: %d\n", col, 0, dir, count)
		//fmt.Println(c)
		if count > bestCount {
			bestCount = count
			bestSq = sq
			bestEntry = dir
		}
	}

	for col := 0; col < len(c[0]); col++ {
		check(0, col, North)
		check(len(c)-1, col, South)
	}
	for row := 0; row < len(c); row++ {
		check(row, 0, West)
		check(row, len(c[0])-1, East)
	}

	return bestSq, bestEntry, bestCount
}

func (c Contraption) EnergisedSquares() int {
	count := 0
	for _, row := range c {
		for _, sq := range row {
			if sq.IsEnergised() {
				count++
			}
		}
	}
	return count
}

func main() {
	c := BuildContraption(helpers.ScanGrid(os.Stdin, ""))
	fmt.Println(c)

	c.ProcessBeam()
	fmt.Println("With beam:")
	//fmt.Println(c)
	fmt.Println("Energised squares:", c.EnergisedSquares())

	c.Reset()
	square, direction, best := c.FindBestBeam()
	fmt.Printf("Most energised squares %d from (%d,%d) %v\n", best, square.Col, square.Row, direction)
}
