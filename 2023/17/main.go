package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Grid [][]int

func (g Grid) String() string {
	var b = strings.Builder{}
	for i, line := range g {
		for _, block := range line {
			b.WriteString(strconv.Itoa(block))
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
		gridLine := make([]int, 0, len(line))
		for _, digit := range strings.Split(line, "") {
			gridLine = append(gridLine, helpers.MustAtoi(digit))
		}
		grid = append(grid, gridLine)
	}
	return grid
}

func (g Grid) Get(pos helpers.Pos) int {
	return g[pos.Y][pos.X]
}

func (g Grid) MaxPos() helpers.Pos {
	return helpers.Pos{X: len(g) - 1, Y: len(g[0]) - 1}
}

type Step struct {
	Pos      helpers.Pos
	Vertical bool
}

type Move struct {
	Step          Step
	TotalHeatLoss int
}

func (m *Move) NextMoves(g Grid, minStep, maxStep int) []*Move {
	maxPos := g.MaxPos()
	var nextMoves []*Move

	// positive direction
	totalHeat := m.TotalHeatLoss
	for i := 1; i <= maxStep; i++ {
		nextPos := m.Step.Pos
		if m.Step.Vertical {
			nextPos.X += i
			if nextPos.X > maxPos.X {
				break
			}
		} else {
			nextPos.Y += i
			if nextPos.Y > maxPos.Y {
				break
			}
		}
		totalHeat += g.Get(nextPos)

		if i >= minStep {
			nextMoves = append(nextMoves, &Move{
				Step:          Step{Pos: nextPos, Vertical: !m.Step.Vertical},
				TotalHeatLoss: totalHeat,
			})
		}
	}
	// Negative direction
	totalHeat = m.TotalHeatLoss
	for i := 1; i <= maxStep; i++ {
		nextPos := m.Step.Pos
		if m.Step.Vertical {
			nextPos.X -= i
			if nextPos.X < 0 {
				break
			}
		} else {
			nextPos.Y -= i
			if nextPos.Y < 0 {
				break
			}
		}
		totalHeat += g.Get(nextPos)

		if i >= minStep {
			nextMoves = append(nextMoves, &Move{
				Step:          Step{Pos: nextPos, Vertical: !m.Step.Vertical},
				TotalHeatLoss: totalHeat,
			})
		}
	}

	return nextMoves
}

type MoveQueue []*Move

func (mq MoveQueue) Len() int           { return len(mq) }
func (mq MoveQueue) Less(i, j int) bool { return mq[i].TotalHeatLoss < mq[j].TotalHeatLoss }
func (mq MoveQueue) Swap(i, j int)      { mq[i], mq[j] = mq[j], mq[i] }

func (mq *MoveQueue) Push(x any) { *mq = append(*mq, x.(*Move)) }
func (mq *MoveQueue) Pop() any {
	old := *mq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*mq = old[0 : n-1]
	return item
}

func BestHeatLoss(g Grid, minStep, maxStep int) int {
	endPos := g.MaxPos()

	seen := make(map[Step]bool)
	moves := &MoveQueue{
		{Step: Step{Vertical: true}}, // Zero Pos and heat
		{Step: Step{Vertical: false}},
	}

	for moves.Len() > 0 {
		move := heap.Pop(moves).(*Move)

		if move.Step.Pos == endPos {
			return move.TotalHeatLoss
		}
		if seen[move.Step] {
			continue
		}
		seen[move.Step] = true

		for _, nextMove := range move.NextMoves(g, minStep, maxStep) {
			if seen[nextMove.Step] {
				continue
			}
			heap.Push(moves, nextMove)
		}
	}
	return -1
}

func main() {
	grid := BuildGrid(helpers.StreamLines(os.Stdin))

	//fmt.Println(grid)

	fmt.Println("Best heat loss:", BestHeatLoss(grid, 1, 3))
	fmt.Println("Best heat loss with ultra:", BestHeatLoss(grid, 4, 10))
}
