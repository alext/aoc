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
	TL   Pos
	BR   Pos
	Area int
}

func NewRect(a, b Pos) Rect {
	return Rect{
		A:    a,
		B:    b,
		TL:   Pos{X: min(a.X, b.X), Y: min(a.Y, b.Y)},
		BR:   Pos{X: max(a.X, b.X), Y: max(a.Y, b.Y)},
		Area: RectArea(a, b),
	}
}

func (p Rect) String() string {
	return fmt.Sprintf("Corners: %s->%s, Area: %d", p.A, p.B, p.Area)
}

type Line struct {
	Start Pos
	End   Pos
}

// Touches returns whether l touches r in any way.
func (l Line) Touches(r Rect) bool {
	if (max(l.Start.X, l.End.X) < r.TL.X) || (min(l.Start.X, l.End.X) > r.BR.X) {
		return false
	}
	if (max(l.Start.Y, l.End.Y) < r.TL.Y) || (min(l.Start.Y, l.End.Y) > r.BR.Y) {
		return false
	}
	return true
}

// GoesInside returns whether l touches any of the inner squares of r not
// including the squares along the edges.
func (l Line) GoesInside(r Rect) bool {
	if (max(l.Start.X, l.End.X) <= r.TL.X) || (min(l.Start.X, l.End.X) >= r.BR.X) {
		return false
	}
	if (max(l.Start.Y, l.End.Y) <= r.TL.Y) || (min(l.Start.Y, l.End.Y) >= r.BR.Y) {
		return false
	}
	return true
}

// ParallelLeft returns a line that is one square to the left of l when
// traversing from Start to End. Assuming a clockwise polygon, this will be one
// square to the outside of the polygon.
func (l Line) ParallelLeft() Line {
	if l.Start.X == l.End.X {
		xDelta := 1
		if l.Start.Y > l.End.Y {
			xDelta = -1
		}
		return Line{
			Start: Pos{X: l.Start.X + xDelta, Y: l.Start.Y},
			End:   Pos{X: l.End.X + xDelta, Y: l.End.Y},
		}

	} else {
		yDelta := 1
		if l.Start.X < l.End.X {
			yDelta = -1
		}
		return Line{
			Start: Pos{X: l.Start.X, Y: l.Start.Y + yDelta},
			End:   Pos{X: l.End.X, Y: l.End.Y + yDelta},
		}

	}
}

type Area struct {
	Lines []Line
}

func BuildArea(points []Pos) *Area {
	a := Area{
		Lines: make([]Line, 0, len(points)),
	}
	for i := 1; i < len(points); i++ {
		start, end := points[i-1], points[i]
		a.Lines = append(a.Lines, Line{Start: start, End: end})
	}
	a.Lines = append(a.Lines, Line{Start: points[len(points)-1], End: points[0]})
	return &a
}

func (a *Area) RectangleInside(r Rect) bool {
	for _, line := range a.Lines {
		if !line.Touches(r) {
			continue
		}
		if line.GoesInside(r) {
			return false
		}
		// line touches edge of r, see if the line to the right (outer side) of
		// r goes inside it or not.
		if line.ParallelLeft().GoesInside(r) {
			return false
		}
	}
	return true
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

	a := BuildArea(redTiles)
	for _, r := range rectangles {
		if a.RectangleInside(r) {
			fmt.Println("Largest inside:", r)
			break
		}
	}
}
