package main

import (
	"flag"
	"fmt"
)

type Elf struct {
	Position int
	Presents int
}

type Circle []*Elf

func NewCircle(numElves int) Circle {
	c := make(Circle, numElves)
	for i := 0; i < len(c); i++ {
		c[i] = &Elf{Position: i + 1, Presents: 1}
	}
	return c
}

func main() {
	numElves := flag.Int("elves", 5, "Number of Elves")
	flag.Parse()

	circle := NewCircle(*numElves)

	lastPos := len(circle) + 1
	lastIndex := 0
	for len(circle) > 1 {
		var elfIndex = -1
		for i := lastIndex; i < len(circle); i++ {
			if circle[i].Position > lastPos {
				elfIndex = i
				break
			}
		}
		if elfIndex == -1 {
			elfIndex = 0
		}
		lastPos = circle[elfIndex].Position

		targetIndex := (len(circle)/2 + elfIndex) % len(circle)

		circle = circle[:targetIndex+copy(circle[targetIndex:], circle[targetIndex+1:])]

		if elfIndex > targetIndex {
			lastIndex = elfIndex - 1
		} else {
			lastIndex = elfIndex
		}
	}

	fmt.Println("Winner at position:", circle[0].Position)
}
