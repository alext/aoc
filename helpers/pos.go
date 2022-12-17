package helpers

import (
	"fmt"
	"strings"
)

type Pos struct {
	X int
	Y int
}

func ParsePos(in string) Pos {
	x, y, _ := strings.Cut(in, ",")
	return Pos{
		X: MustAtoi(x),
		Y: MustAtoi(y),
	}
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// DistanceTo returns the Manhatten distance between p and other.
func (p Pos) DistanceTo(other Pos) int {
	return AbsInt(p.X-other.X) + AbsInt(p.Y-other.Y)
}
