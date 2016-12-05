package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

type House struct {
	X int
	Y int
}

func main() {
	deliveries := make(map[House]int)
	santaPos := House{0, 0}
	roboPos := House{0, 0}
	deliveries[santaPos] = 2
	santasTurn := true

	helpers.ScanRunes(os.Stdin, func(t string) {
		var current House
		if santasTurn {
			current = santaPos
		} else {
			current = roboPos
		}
		switch t {
		case ">":
			current.X += 1
		case "<":
			current.X -= 1
		case "^":
			current.Y += 1
		case "v":
			current.Y -= 1
		default:
			fmt.Println("Unexpected char in input:", t)
			return
		}
		if _, ok := deliveries[current]; ok {
			deliveries[current] += 1
		} else {
			deliveries[current] = 1
		}
		if santasTurn {
			santaPos = current
		} else {
			roboPos = current
		}
		santasTurn = !santasTurn
	})

	fmt.Println("Houses visited:", len(deliveries))
}
