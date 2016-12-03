package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction uint8

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) Turn(dir string) Direction {
	if dir == "L" {
		if d == North {
			return West
		} else {
			return d - 1
		}
	} else {
		if d == West {
			return North
		} else {
			return d + 1
		}
	}
}

type Position struct {
	x int
	y int
}

type Navigator struct {
	pos     Position
	History map[Position]bool
	d       Direction
}

func NewNavigator() *Navigator {
	return &Navigator{
		History: make(map[Position]bool),
	}
}

func (n *Navigator) visited(p Position) bool {
	_, result := n.History[p]
	n.History[p] = true
	return result
}

func (n *Navigator) Move(turn string, steps int) bool {
	n.d = n.d.Turn(turn)
	newpos := n.pos
	for i := 0; i < steps; i++ {
		switch n.d {
		case North:
			newpos.y += 1
		case East:
			newpos.x += 1
		case South:
			newpos.y -= 1
		case West:
			newpos.x -= 1
		}
		if n.visited(newpos) {
			n.pos = newpos
			return true
		}
	}
	n.pos = newpos
	return false
}

func (n *Navigator) Distance() int {
	return absInt(n.pos.x) + absInt(n.pos.y)
}

func (n *Navigator) String() string {
	return fmt.Sprintf("x: %d, y: %d, Distance: %d", n.pos.x, n.pos.y, n.Distance())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if atEOF {
			return len(data), data, bufio.ErrFinalToken
		} else {
			return 0, nil, nil
		}
	})

	n := NewNavigator()
	var turn string
	var steps int

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%1s%d", &turn, &steps)
		if n.Move(turn, steps) {
			fmt.Println("First duplicate pos:", n)
			return
		}
		fmt.Println("Step:", n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Final pos:", n)
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
