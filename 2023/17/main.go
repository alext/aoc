package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Direction uint8

const (
	North Direction = iota + 1
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "^"
	case South:
		return "v"
	case East:
		return ">"
	case West:
		return "<"
	default:
		return "?"
	}
}

type Block struct {
	HeatLoss int
	Row      int
	Col      int

	BestHeatLoss     int
	EntryDirection   Direction
	ConsecutiveCount int
}

func (b Block) Char() string {
	if b.BestHeatLoss > 0 {
		return b.EntryDirection.String()
	}
	return strconv.Itoa(b.HeatLoss)
}

func (b Block) String() string {
	return fmt.Sprintf("(%d,%d) %s %d", b.Col, b.Row, b.EntryDirection, b.BestHeatLoss)
}

type PotentialMove struct {
	Target    *Block
	Direction Direction
}

func (b *Block) NextMoves(g Grid) []*PotentialMove {
	var nextMoves []*PotentialMove
	// Going North
	if b.Row > 0 && b.EntryDirection != South {
		//if b.EntryDirection != North || b.ConsecutiveCount < 3 {
		nextMoves = append(nextMoves, &PotentialMove{
			Target:    g[b.Row-1][b.Col],
			Direction: North,
		})
		//}
	}
	// Going East
	if b.Col < len(g[0])-1 && b.EntryDirection != West {
		//if b.EntryDirection != East || b.ConsecutiveCount < 3 {
		nextMoves = append(nextMoves, &PotentialMove{
			Target:    g[b.Row][b.Col+1],
			Direction: East,
		})
		//}
	}
	// Going South
	if b.Row < len(g)-1 && b.EntryDirection != North {
		//if b.EntryDirection != South || b.ConsecutiveCount < 3 {
		nextMoves = append(nextMoves, &PotentialMove{
			Target:    g[b.Row+1][b.Col],
			Direction: South,
		})
		//}
	}
	// Going West
	if b.Col > 0 && b.EntryDirection != East {
		//if b.EntryDirection != West || b.ConsecutiveCount < 3 {
		nextMoves = append(nextMoves, &PotentialMove{
			Target:    g[b.Row][b.Col-1],
			Direction: West,
		})
		//}
	}
	return nextMoves
}

type Grid [][]*Block

func (g Grid) String() string {
	var b = strings.Builder{}
	for i, line := range g {
		for _, block := range line {
			b.WriteString(block.Char())
		}
		if i < len(g)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func BuildGrid(input <-chan string) Grid {
	var grid Grid
	for line := range input {
		gridLine := make([]*Block, 0, len(line))
		for i, digit := range strings.Split(line, "") {
			gridLine = append(gridLine, &Block{
				HeatLoss: helpers.MustAtoi(digit),
				Row:      len(grid),
				Col:      i,
			})

		}
		grid = append(grid, gridLine)
	}
	return grid
}

func (g Grid) BestPath() int {
	start := g[0][0]
	//start.EntryDirection = West
	//start.ConsecutiveCount = 3
	moves := []*Block{start}

	limit := 200
	for len(moves) > 0 && limit > 0 {
		move := moves[0]
		moves = moves[1:]

		if move.Row == len(g)-1 && move.Col == len(g[0])-1 {
			// We've reached the final square
			return move.BestHeatLoss
		}

		for _, potentialMove := range move.NextMoves(g) {
			target := potentialMove.Target
			if move.EntryDirection == potentialMove.Direction && move.ConsecutiveCount >= 3 {
				// Can't move more than 3 in a row
				continue
			}
			potentialBest := move.BestHeatLoss + target.HeatLoss
			if target.BestHeatLoss != 0 && target.BestHeatLoss < potentialBest {
				// Already a better path to target
				continue
			}
			target.BestHeatLoss = potentialBest
			target.EntryDirection = potentialMove.Direction
			if potentialMove.Direction == move.EntryDirection {
				target.ConsecutiveCount = move.ConsecutiveCount + 1
			} else {
				target.ConsecutiveCount = 1
			}
			if !slices.Contains(moves, target) {
				moves = append(moves, target)
			}
		}

		slices.SortFunc(moves, func(a, b *Block) int { return cmp.Compare(a.BestHeatLoss, b.BestHeatLoss) })
		//fmt.Println("After move:", moves)
		//limit--
	}

	return -1
}

func main() {
	g := BuildGrid(helpers.StreamLines(os.Stdin))

	fmt.Println(g)

	fmt.Println("Heat loss on best path:", g.BestPath())

	fmt.Println(g)
}
