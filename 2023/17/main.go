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
	Pos      helpers.Pos
}

func (b Block) String() string {
	return strconv.Itoa(b.HeatLoss)
}

type Grid [][]*Block

func (g Grid) String() string {
	var b = strings.Builder{}
	for i, line := range g {
		for _, block := range line {
			b.WriteString(block.String())
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
				Pos:      helpers.Pos{X: i, Y: len(grid)},
			})

		}
		grid = append(grid, gridLine)
	}
	return grid
}

type Move struct {
	Path           []*Block
	Target         *Block
	EntryDirection Direction

	DistanceToEnd    int
	TotalHeatLoss    int
	ConsecutiveCount int
	Beaten           bool
}

func (m *Move) String() string {
	return fmt.Sprintf("%d%s%s%d", m.TotalHeatLoss, m.Target.Pos, m.EntryDirection, m.ConsecutiveCount)
}

func (m *Move) PathStr() string {
	var b strings.Builder
	for _, blk := range m.Path {
		b.WriteString(blk.Pos.String())
		b.WriteString(" ")
	}
	b.WriteString(m.Target.Pos.String())
	return b.String()
}

func (m *Move) NextMoves(g Grid) []*Move {
	var nextMoves []*Move
	b := m.Target
	nextPath := append(slices.Clone(m.Path), b)
	// Going North
	if b.Pos.Y > 0 && m.EntryDirection != South {
		target := g[b.Pos.Y-1][b.Pos.X]
		if !slices.Contains(m.Path, target) {
			nextMoves = append(nextMoves, &Move{
				Path:           nextPath,
				Target:         target,
				EntryDirection: North,
			})
		}
	}
	// Going East
	if b.Pos.X < len(g[0])-1 && m.EntryDirection != West {
		target := g[b.Pos.Y][b.Pos.X+1]
		if !slices.Contains(m.Path, target) {
			nextMoves = append(nextMoves, &Move{
				Path:           nextPath,
				Target:         target,
				EntryDirection: East,
			})
		}
	}
	// Going South
	if b.Pos.Y < len(g)-1 && m.EntryDirection != North {
		target := g[b.Pos.Y+1][b.Pos.X]
		if !slices.Contains(m.Path, target) {
			nextMoves = append(nextMoves, &Move{
				Path:           nextPath,
				Target:         target,
				EntryDirection: South,
			})
		}
	}
	// Going West
	if b.Pos.X > 0 && m.EntryDirection != East {
		target := g[b.Pos.Y][b.Pos.X-1]
		if !slices.Contains(m.Path, target) {
			nextMoves = append(nextMoves, &Move{
				Path:           nextPath,
				Target:         target,
				EntryDirection: West,
			})
		}
	}
	return nextMoves
}

// Returns whether m is guaranteed to be better than other based on the move
// properties alone.
func (m *Move) BetterThan(other *Move) bool {
	if m.TotalHeatLoss < other.TotalHeatLoss && m.EntryDirection == other.EntryDirection && m.ConsecutiveCount <= other.ConsecutiveCount {
		return true
	}
	return false
}

func (g Grid) BestPath() *Move {
	start := &Move{
		Target: g[0][0],
	}
	endBlock := g[len(g)-1][len(g[0])-1]
	moves := []*Move{start}

	movesToBlock := make(map[*Block][]*Move)

	limit := 3000
	for len(moves) > 0 && limit > 0 {
		move := moves[0]
		moves = moves[1:]

		if move.Target == endBlock {
			return move
			//if best == nil || move.TotalHeatLoss < best.TotalHeatLoss {
			//    fmt.Printf("Found new best loss=%d, path: %s\n", move.TotalHeatLoss, move.PathStr())
			//    best = move
			//    continue
			//}
		}

		if move.Beaten {
			continue
		}

		// best could have improved since this was added to the list
		//if best != nil && (move.TotalHeatLoss+move.DistanceToEnd) >= best.TotalHeatLoss {
		//    continue
		//}

		for _, nextMove := range move.NextMoves(g) {
			target := nextMove.Target
			if move.EntryDirection == nextMove.EntryDirection && move.ConsecutiveCount >= 3 {
				// Can't move more than 3 in a row
				continue
			}
			nextMove.TotalHeatLoss = move.TotalHeatLoss + target.HeatLoss
			//nextMove.DistanceToEnd = nextMove.Target.Pos.DistanceTo(endBlock.Pos)

			//if best != nil && (nextMove.TotalHeatLoss+move.DistanceToEnd) >= best.TotalHeatLoss {
			//    continue
			//}

			if nextMove.EntryDirection == move.EntryDirection {
				nextMove.ConsecutiveCount = move.ConsecutiveCount + 1
			} else {
				nextMove.ConsecutiveCount = 1
			}
			for _, otherMove := range movesToBlock[target] {
				if otherMove.BetterThan(nextMove) {
					nextMove.Beaten = true
					break
				}
				if nextMove.BetterThan(otherMove) {
					otherMove.Beaten = true
					// TODO: remove from movesToBlock
				}
			}
			movesToBlock[target] = slices.DeleteFunc(movesToBlock[target], func(m *Move) bool { return m.Beaten })
			if nextMove.Beaten {
				continue
			}

			moves = append(moves, nextMove)
			movesToBlock[target] = append(movesToBlock[target], nextMove)
		}

		slices.SortFunc(moves, func(a, b *Move) int { return cmp.Compare(a.TotalHeatLoss, b.TotalHeatLoss) })
		//fmt.Println("After move:", moves)
		//limit--
	}

	//if best != nil {
	//    return best.TotalHeatLoss
	//}
	return nil
}

func main() {
	g := BuildGrid(helpers.StreamLines(os.Stdin))

	//fmt.Println(g)

	move := g.BestPath()
	fmt.Println("Best path:", move, move.PathStr())
	fmt.Println("Heat loss on best path:", move.TotalHeatLoss)
}
