package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type House struct {
	X int
	Y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	deliveries := make(map[House]int)
	santaPos := House{0, 0}
	roboPos := House{0, 0}
	deliveries[santaPos] = 2
	santasTurn := true

	for scanner.Scan() {
		var current House
		if santasTurn {
			current = santaPos
		} else {
			current = roboPos
		}
		switch t := scanner.Text(); t {
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
			continue
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
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Houses visited:", len(deliveries))
}
