package main

import (
	"flag"
	"fmt"
)

type Circle []int

func NewCircle(numElves int) Circle {
	c := make(Circle, numElves)
	for i := 0; i < len(c); i++ {
		c[i] = 1
	}
	return c
}

func (c Circle) playRound() {
	length := len(c)
	for i := 0; i < length; i++ {
		if c[i] == 0 {
			continue
		}
		for j := (i + 1) % length; j != i; j = (j + 1) % length {
			if c[j] > 0 {
				c[i] += c[j]
				c[j] = 0
				break
			}
		}
	}
}

func (c Circle) Play() int {
	for {
		c.playRound()
		for i := 1; i < len(c); i++ {
			if c[i] == len(c) {
				return i + 1
			}
		}
	}
	return 0
}

func main() {
	numElves := flag.Int("elves", 5, "Number of Elves")
	flag.Parse()

	circle := NewCircle(*numElves)
	winner := circle.Play()

	fmt.Println("Winner at position:", winner)
}
