package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

type Pos = helpers.Pos

type Plot struct {
	Pos      Pos
	MinSteps int
}

func (p Plot) String() string {
	return fmt.Sprintf("%s:%d", p.Pos, p.MinSteps)
}

type Garden struct {
	Plots     map[Pos]*Plot
	Start     *Plot
	MaxExtent Pos
	GridSize  int
}

func ParseGarden(input [][]string) *Garden {
	if len(input) != len(input[0]) {
		log.Fatal("Non-square map...")
	}
	g := &Garden{
		Plots:     make(map[Pos]*Plot),
		MaxExtent: Pos{Y: len(input) - 1, X: len(input[0]) - 1},
		GridSize:  len(input),
	}
	for y, line := range input {
		for x, ch := range line {
			if ch == "#" {
				continue
			}
			plot := &Plot{
				Pos: Pos{X: x, Y: y},
			}
			if ch == "S" {
				g.Start = plot
			}
			g.Plots[plot.Pos] = plot
		}
	}
	return g
}

func roundPos(pos Pos, gridSize int) Pos {
	// Double modulus to handle negative inputs
	return Pos{
		X: (pos.X%gridSize + gridSize) % gridSize,
		Y: (pos.Y%gridSize + gridSize) % gridSize,
	}
}

func (g *Garden) GetPlot(pos Pos) *Plot {
	if pos.X >= 0 && pos.Y >= 0 && pos.X <= g.MaxExtent.X && pos.Y <= g.MaxExtent.Y {
		// Within center grid - just return
		return g.Plots[pos]
	}
	if existing, ok := g.Plots[pos]; ok {
		// One we've already calculated
		return existing
	}
	roundedPos := roundPos(pos, g.GridSize)
	if _, ok := g.Plots[roundedPos]; !ok {
		// Rock in initial grid
		return nil
	}
	plot := &Plot{Pos: pos}
	g.Plots[pos] = plot
	return plot
}

func (g *Garden) CountPlots(stepCount int) int {
	current := map[Pos]*Plot{g.Start.Pos: g.Start}
	next := make(map[Pos]*Plot)

	seen := map[Pos]*Plot{g.Start.Pos: g.Start}
	for step := 1; step <= stepCount && len(current) > 0; step++ {
		for _, plot := range current {
			for _, pos := range plot.Pos.Neighbours() {
				if _, seen := seen[pos]; seen {
					continue
				}
				nextPlot := g.GetPlot(pos)
				if nextPlot == nil {
					continue
				}
				nextPlot.MinSteps = step
				next[pos] = nextPlot
				seen[pos] = nextPlot
			}
		}
		current, next = next, current
		clear(next)
	}
	count := 0
	oddEven := stepCount % 2
	for _, plot := range seen {
		if plot.MinSteps <= stepCount && plot.MinSteps%2 == oddEven {
			count++
		}
	}
	return count
}

func getIntervals(input []int, depth int) []int {
	//fmt.Println(input)
	var intervals []int
	for i := 1; i < len(input); i++ {
		intervals = append(intervals, input[i]-input[i-1])
	}
	if depth <= 1 {
		return intervals
	}
	return getIntervals(intervals, depth-1)
}

func calculateTotalPlots(initial, initialDelta, numGrids int, delta2ndOrder int) int {
	//fmt.Println("calculateTotalPlots", initial, initialDelta, numGrids, delta2ndOrder)
	plotCount := initial
	nextDelta := initialDelta
	for gridCount := 1; gridCount < numGrids; gridCount++ {
		plotCount += nextDelta
		nextDelta += delta2ndOrder
	}
	return plotCount
}

func main() {
	garden := ParseGarden(helpers.ScanGrid(os.Stdin, ""))
	//fmt.Println(garden)

	if garden.GridSize < 15 {
		// Small input
		fmt.Println("Plots within 6 steps:", garden.CountPlots(6))
	} else {
		// real input
		fmt.Println("Plots within 64 steps:", garden.CountPlots(64))
	}

	// Part 2
	if garden.GridSize < 15 {
		for _, n := range []int{6, 10, 50, 100} {
			fmt.Printf("Plots with %d steps: %d\n", n, garden.CountPlots(n))
		}
		return
	}

	const stepCount = 26501365
	numGrids := stepCount / garden.GridSize
	remainder := stepCount % garden.GridSize
	fmt.Printf("%d Steps is %d squares with %d remaining\n", stepCount, numGrids, remainder)
	var sequence []int
	var deltas []int
	var deltas2ndOrder []int
	for n := remainder + garden.GridSize; n < 1000; n += garden.GridSize {
		numPlots := garden.CountPlots(n)
		fmt.Printf("Plots with %d steps: %d\n", n, numPlots)
		sequence = append(sequence, numPlots)
		if len(sequence) == 3 {
			deltas = getIntervals(sequence, 1)
			deltas2ndOrder = getIntervals(sequence, 2)
			fmt.Println("Sequence", sequence)
			fmt.Println("1st order intervals", deltas)
			fmt.Println("2nd order intervals", deltas2ndOrder)
			break
		}
	}

	//fmt.Println("Total plots 2", calculateTotalPlots(sequence[0], deltas[0], 2, deltas2ndOrder[0]))
	//fmt.Println("Total plots 3", calculateTotalPlots(sequence[0], deltas[0], 3, deltas2ndOrder[0]))
	//fmt.Println("Total plots 4", calculateTotalPlots(sequence[0], deltas[0], 4, deltas2ndOrder[0]))
	//fmt.Println("Total plots 5", calculateTotalPlots(sequence[0], deltas[0], 5, deltas2ndOrder[0]))
	//fmt.Println("Total plots 6", calculateTotalPlots(sequence[0], deltas[0], 6, deltas2ndOrder[0]))
	totalPlots := calculateTotalPlots(sequence[0], deltas[0], numGrids, deltas2ndOrder[0])
	fmt.Println("Total plots", totalPlots)
}
