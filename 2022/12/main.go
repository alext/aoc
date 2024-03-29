package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

type Square struct {
	Height   int
	Row      int
	Col      int
	Distance int
}

func NewSquare(height, row, col int) *Square {
	return &Square{
		Height:   height,
		Row:      row,
		Col:      col,
		Distance: -1,
	}
}

func (s Square) String() string {
	return fmt.Sprintf("%s(%d,%d)", string('a'+s.Height), s.Col, s.Row)
}

func (s Square) Visited() bool {
	return s.Distance >= 0
}

func (s *Square) getMoves(m *Map, test func(*Square) bool) []*Square {
	var moves []*Square
	for _, neighbour := range []*Square{
		m.GetSquare(s.Row+1, s.Col),
		m.GetSquare(s.Row-1, s.Col),
		m.GetSquare(s.Row, s.Col+1),
		m.GetSquare(s.Row, s.Col-1),
	} {
		if neighbour == nil {
			continue
		}
		if test(neighbour) {
			moves = append(moves, neighbour)
		}
	}
	return moves
}

func (s *Square) PossibleMoves(m *Map) []*Square {
	return s.getMoves(
		m,
		func(neighbour *Square) bool {
			return neighbour.Height-s.Height <= 1
		},
	)
}

func (s *Square) PossibleMovesReverse(m *Map) []*Square {
	return s.getMoves(
		m,
		func(neighbour *Square) bool {
			return s.Height-neighbour.Height <= 1
		},
	)
}

type Map struct {
	Squares [][]*Square
	Start   *Square
	End     *Square
}

func (m *Map) Reset() {
	for _, row := range m.Squares {
		for _, s := range row {
			s.Distance = -1
		}
	}
}

func (m *Map) GetSquare(row, col int) *Square {
	if row < 0 || row >= len(m.Squares) {
		return nil
	}
	if col < 0 || col >= len(m.Squares[row]) {
		return nil
	}
	return m.Squares[row][col]
}

func (m *Map) findPath(start *Square, getMoves func(*Square) []*Square, test func(*Square) bool) int {
	var currentList, nextList []*Square
	start.Distance = 0
	currentList = append(currentList, start)

	for len(currentList) > 0 {
		for _, current := range currentList {
			for _, neighbour := range getMoves(current) {
				if neighbour.Visited() {
					continue
				}
				neighbour.Distance = current.Distance + 1
				if test(neighbour) {
					return neighbour.Distance
				}
				nextList = append(nextList, neighbour)
			}
		}
		currentList = nextList
		nextList = nil
	}

	log.Fatal("Failed to find a path")
	return 0
}

func (m *Map) ShortestPath() int {
	return m.findPath(
		m.Start,
		func(s *Square) []*Square { return s.PossibleMoves(m) },
		func(s *Square) bool { return s == m.End },
	)
}

func (m *Map) ShortestPathAny() int {
	return m.findPath(
		m.End,
		func(s *Square) []*Square { return s.PossibleMovesReverse(m) },
		func(s *Square) bool { return s.Height == 0 },
	)
}

func main() {
	var m Map

	row := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		mapLine := make([]*Square, 0, len(line))
		for col, chr := range line {
			switch chr {
			case 'S':
				sq := NewSquare(0, row, col)
				mapLine = append(mapLine, sq)
				m.Start = sq
			case 'E':
				sq := NewSquare('z'-'a', row, col)
				mapLine = append(mapLine, sq)
				m.End = sq
			default:
				mapLine = append(mapLine, NewSquare(int(chr-'a'), row, col))
			}
		}
		m.Squares = append(m.Squares, mapLine)
		row++
	})

	fmt.Println("Shortest distance", m.ShortestPath())

	m.Reset()

	fmt.Println("Shortest distance from any a", m.ShortestPathAny())
}
