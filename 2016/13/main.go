package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

var seed uint64 = 1364

type Position struct {
	X uint64
	Y uint64
}

func (p *Position) IsWall() bool {
	val := p.X*p.X + 3*p.X + 2*p.X*p.Y + p.Y + p.Y*p.Y + seed
	binary := strconv.FormatUint(val, 2)
	numBits := strings.Count(binary, "1")
	return numBits%2 != 0
}

func (p Position) Neighbours() <-chan Position {
	ch := make(chan Position, 2)
	go func() {
		if p.X > 0 {
			ch <- Position{X: p.X - 1, Y: p.Y}
		}
		ch <- Position{X: p.X + 1, Y: p.Y}
		if p.Y > 0 {
			ch <- Position{X: p.X, Y: p.Y - 1}
		}
		ch <- Position{X: p.X, Y: p.Y + 1}

		close(ch)
	}()
	return ch
}

type Space struct {
	Wall     bool
	Pos      Position
	Distance uint64
	Visited  bool
}

// Represents an as-yet unknown distance
const Infinity = 999999

func NewSpace(p Position) *Space {
	return &Space{
		Pos:      p,
		Distance: Infinity,
		Wall:     p.IsWall(),
	}
}

type Maze map[Position]*Space

func (m Maze) iterateSpaces(ctx context.Context, start *Space) <-chan *Space {
	start.Distance = 0
	m[start.Pos] = start

	candidates := make(map[Position]bool)
	candidates[start.Pos] = true

	ch := make(chan *Space)
	go func() {
		for {
			if len(candidates) == 0 {
				close(ch)
				return
			}

			var (
				distance  uint64 = Infinity
				candidate *Space
			)
			for pos, _ := range candidates {
				if m[pos].Distance < distance {
					candidate = m[pos]
					distance = candidate.Distance
				}
			}
			delete(candidates, candidate.Pos)

			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- candidate:
			}

			for pos := range candidate.Pos.Neighbours() {
				space, ok := m[pos]
				if !ok {
					space = NewSpace(pos)
					m[pos] = space
				}
				if space.Wall || space.Visited {
					continue
				}
				if space.Distance > candidate.Distance+1 {
					space.Distance = candidate.Distance + 1
					candidates[space.Pos] = true
				}
			}

			candidate.Visited = true
		}
	}()
	return ch
}

func (m Maze) ShortestPath(startPos, targetPos Position) uint64 {
	target := NewSpace(targetPos)
	if target.Wall {
		log.Fatalln("Error target square is a wall")
	}
	m[target.Pos] = target

	start := NewSpace(startPos)
	if start.Wall {
		log.Fatalln("Error start square is a wall")
	}

	ctx, cancel := context.WithCancel(context.Background())
	for space := range m.iterateSpaces(ctx, start) {
		if space.Pos == targetPos {
			cancel()
			return space.Distance
		}
	}
	log.Fatalln("No more candidate spaces")
	return Infinity
}

func (m Maze) ReachableSpaces(startPos Position, maxDistance uint64) int {
	start := NewSpace(startPos)
	if start.Wall {
		log.Fatalln("Error start square is a wall")
	}

	maxDistance -= 1 // Starting square counts as a step apparently

	count := 1
	ctx, cancel := context.WithCancel(context.Background())
	for space := range m.iterateSpaces(ctx, start) {
		if space.Distance > maxDistance {
			cancel()
			return count
		}
		count++
	}
	// All reachable spaces covered
	return count
}

func main() {
	targetFlag := flag.String("target", "", "Target position")
	maxDistanceFlag := flag.Int("max-distance", 0, "Count locations reachable within N steps")
	flag.Parse()

	if *targetFlag != "" {
		targetPos := targetPosition(*targetFlag)

		m := make(Maze)
		distance := m.ShortestPath(Position{X: 1, Y: 1}, targetPos)
		fmt.Println("Shortest path:", distance)
	}

	maxDistance := uint64(*maxDistanceFlag)
	if maxDistance > 0 {
		m := make(Maze)
		count := m.ReachableSpaces(Position{X: 1, Y: 1}, maxDistance)
		fmt.Println("Locations reachable in", maxDistance, "steps or fewer:", count)
	}
}

func targetPosition(param string) Position {
	parts := strings.SplitN(param, ",", 2)
	return Position{
		X: uint64(helpers.MustAtoi(parts[0])),
		Y: uint64(helpers.MustAtoi(parts[1])),
	}
}
