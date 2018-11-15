package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

func processGroup(input <-chan string, depth uint) (uint, uint) {
	score := depth
	var garbageCount uint = 0
	for r := range input {
		switch r {
		case "}":
			return score, garbageCount
		case "{":
			s, g := processGroup(input, depth+1)
			score += s
			garbageCount += g
		case "<":
			garbageCount += processGarbage(input)
		}
	}
	return score, garbageCount // will never happen
}

func processGarbage(input <-chan string) uint {
	var count uint = 0
	for r := range input {
		switch r {
		case ">":
			return count
		case "!":
			_ = <-input
		default:
			count++
		}
	}
	return count // will never happen
}

func main() {
	input := helpers.StreamRunes(os.Stdin)
	r := <-input
	if r != "{" {
		log.Fatalf("Unexpected char %s at start of input", r)
	}
	score, garbageCount := processGroup(input, 1)
	fmt.Println("Total score", score)
	fmt.Println("Total garbage", garbageCount)
}
