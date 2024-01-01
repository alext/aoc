package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Pos3 struct {
	X, Y, Z int
}

func (p Pos3) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func MaxPos(a, b Pos3) Pos3 {
	return Pos3{
		X: max(a.X, b.X),
		Y: max(a.Y, b.Y),
		Z: max(a.Z, b.Z),
	}
}

type Brick struct {
	Pos1  Pos3
	Pos2  Pos3
	Label string
}

func (b Brick) String() string {
	return fmt.Sprintf("%s-%s", b.Pos1, b.Pos2)
}

func ParseBrick(line string) *Brick {
	b := Brick{}
	// 1,0,1~1,2,1
	_, err := fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &b.Pos1.X, &b.Pos1.Y, &b.Pos1.Z, &b.Pos2.X, &b.Pos2.Y, &b.Pos2.Z)
	if err != nil {
		log.Fatalln("Failed to parse line", line, err)
	}
	if b.Pos1.X > b.Pos2.X || b.Pos1.Y > b.Pos2.Y || b.Pos1.Z > b.Pos2.Z {
		log.Fatalln("Positions out of order in", line)
	}
	return &b
}

func (b Brick) Positions() []Pos3 {
	var positions []Pos3
	for x := b.Pos1.X; x <= b.Pos2.X; x++ {
		for y := b.Pos1.Y; y <= b.Pos2.Y; y++ {
			for z := b.Pos1.Z; z <= b.Pos2.Z; z++ {
				positions = append(positions, Pos3{X: x, Y: y, Z: z})
			}
		}
	}
	return positions
}
func (b Brick) PositionsBelow() []Pos3 {
	var positions []Pos3
	for x := b.Pos1.X; x <= b.Pos2.X; x++ {
		for y := b.Pos1.Y; y <= b.Pos2.Y; y++ {
			positions = append(positions, Pos3{X: x, Y: y, Z: b.Pos1.Z - 1})
		}
	}
	return positions
}

type Stack struct {
	Bricks []*Brick
	Grid   map[Pos3]*Brick
	MaxPos Pos3
}

func NewStack() *Stack {
	return &Stack{
		Grid: make(map[Pos3]*Brick),
	}
}

func (s *Stack) String() string {
	var b strings.Builder
	bricksAt := func(x, z int) []*Brick {
		var bricks []*Brick
		for y := 0; y <= s.MaxPos.Y; y++ {
			if b := s.Grid[Pos3{X: x, Y: y, Z: z}]; b != nil {
				bricks = append(bricks, b)
			}
		}
		bricks = slices.Compact(bricks)
		return bricks
	}
	for z := s.MaxPos.Z; z > 0; z-- {
		for x := 0; x <= s.MaxPos.X; x++ {
			bricks := bricksAt(x, z)
			if len(bricks) == 0 {
				b.WriteString(".")
			} else if len(bricks) == 1 {
				b.WriteString(bricks[0].Label)
			} else {
				b.WriteString(strconv.Itoa(len(bricks)))
			}
		}
		b.WriteString("\n")
	}
	b.WriteString(strings.Repeat("-", s.MaxPos.X+1))
	return b.String()
}

func (s *Stack) AddBrick(b *Brick) {
	b.Label = string('A' + len(s.Bricks))
	s.Bricks = append(s.Bricks, b)
	for _, pos := range b.Positions() {
		if s.Grid[pos] != nil {
			log.Fatalln("Multiple bricks at position:", pos)
		}
		s.Grid[pos] = b
		s.MaxPos = MaxPos(s.MaxPos, b.Pos2)
	}
}

func (s *Stack) MoveDown(b *Brick) {
	for _, pos := range b.Positions() {
		delete(s.Grid, pos)
	}
	b.Pos1.Z--
	b.Pos2.Z--
	for _, pos := range b.Positions() {
		if s.Grid[pos] != nil {
			log.Fatalln("Multiple bricks at position:", pos)
		}
		s.Grid[pos] = b
	}
}

func (s *Stack) Settle() {
	slices.SortFunc(s.Bricks, func(a, b *Brick) int { return cmp.Compare(a.Pos1.Z, b.Pos1.Z) })
	for i := 0; i < len(s.Bricks); i++ {
		b := s.Bricks[i]
		if b.Pos1.Z == 1 {
			// On floor...
			continue
		}

		allClear := true
		for _, pos := range b.PositionsBelow() {
			if s.Grid[pos] != nil {
				allClear = false
				break
			}
		}
		if !allClear {
			continue
		}
		s.MoveDown(b)
		i-- // Brick might be able to move down multiple, so re-evaluate
	}
	slices.SortFunc(s.Bricks, func(a, b *Brick) int { return cmp.Compare(a.Pos1.Z, b.Pos1.Z) })
}

func (s *Stack) ChainReactionSize(removedBrick *Brick) int {
	movedBricks := map[*Brick]bool{removedBrick: true}
	for i := slices.Index(s.Bricks, removedBrick) + 1; i < len(s.Bricks); i++ {
		brick := s.Bricks[i]
		if brick.Pos1.Z == 1 {
			// On floor...
			continue
		}
		willFall := true
		for _, pos := range brick.PositionsBelow() {
			b := s.Grid[pos]
			if b != nil && !movedBricks[b] {
				willFall = false
				break
			}
		}
		if !willFall {
			continue
		}
		movedBricks[brick] = true
	}
	return len(movedBricks) - 1 // -1 so as to not include initial disingegrated one
}

func main() {
	s := NewStack()
	helpers.ScanLines(os.Stdin, func(line string) {
		s.AddBrick(ParseBrick(line))
	})
	if len(s.Bricks) < 26 {
		fmt.Println(s)
	}
	s.Settle()
	if len(s.Bricks) < 26 {
		fmt.Println(s)
	}

	stableCount := 0
	totalChainReactionSize := 0
	for _, brick := range s.Bricks {
		count := s.ChainReactionSize(brick)
		totalChainReactionSize += count
		if count == 0 {
			stableCount++
		}
	}
	fmt.Println("Stable bricks:", stableCount)
	fmt.Println("Total chain reaction size:", totalChainReactionSize)
}
