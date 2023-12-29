package main

import (
	"fmt"
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
}

func ParseGarden(input [][]string) *Garden {
	g := &Garden{
		Plots:     make(map[Pos]*Plot),
		MaxExtent: Pos{Y: len(input), X: len(input[0])},
	}
	for y, line := range input {
		for x, ch := range line {
			if ch == "#" {
				continue
			}
			plot := &Plot{
				Pos: Pos{X: x + 1, Y: y + 1},
			}
			if ch == "S" {
				g.Start = plot
			}
			g.Plots[plot.Pos] = plot
		}
	}
	return g
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
				nextPlot, found := g.Plots[pos]
				if !found {
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

func main() {
	garden := ParseGarden(helpers.ScanGrid(os.Stdin, ""))
	fmt.Println(garden)

	if garden.MaxExtent.X < 15 {
		// Small input
		fmt.Println("Plots within 6 steps:", garden.CountPlots(6))
	} else {
		// real input
		fmt.Println("Plots within 64 steps:", garden.CountPlots(64))
	}
}
