package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"

	"github.com/alext/aoc/helpers"
)

type Pos = helpers.Pos

func RectArea(a, b Pos) int {
	return (helpers.AbsInt(a.X-b.X) + 1) * (helpers.AbsInt(a.Y-b.Y) + 1)
}

type Rect struct {
	A    Pos
	B    Pos
	Area int
}

func NewRect(a, b Pos) Rect {
	return Rect{
		A:    a,
		B:    b,
		Area: RectArea(a, b),
	}
}

func (p Rect) String() string {
	return fmt.Sprintf("Corners: %s->%s, Area: %d", p.A, p.B, p.Area)
}

func main() {
	var redTiles []Pos
	helpers.ScanLines(os.Stdin, func(line string) {
		redTiles = append(redTiles, helpers.ParsePos(line))
	})

	var rectangles []Rect
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			rectangles = append(rectangles, NewRect(redTiles[i], redTiles[j]))
		}
	}
	slices.SortFunc(rectangles, func(a, b Rect) int {
		return cmp.Compare(b.Area, a.Area)
	})

	fmt.Println("Largest:", rectangles[0])
}
