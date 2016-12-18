package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alext/aoc/helpers"
)

const floors = 4

type State interface {
	Hash() uint32
	Complete() bool
	Safe() bool
	AvailableMoves() <-chan State
}

var (
	floorRE = regexp.MustCompile(`The (\S+) floor contains`)
	itemRE  = regexp.MustCompile(`an? (.*?)(?:-compatible)? (generator|microchip)`)
)

func readFloor(line string) int {
	matches := floorRE.FindStringSubmatch(line)
	if matches == nil {
		log.Fatalln("Error reading floor number in line:", line)
	}
	switch matches[1] {
	case "first":
		return 0
	case "second":
		return 1
	case "third":
		return 2
	case "fourth":
		return 3
	default:
		log.Fatalln("Unrecognised floor number:", matches[1])
	}
	return 0 // Never reached
}

func readItems(line string) []string {
	allMatches := itemRE.FindAllStringSubmatch(line, -1)
	if allMatches == nil {
		log.Fatalln("Failed to find items for line:", line)
	}
	var results []string
	for _, matches := range allMatches {
		results = append(results, matches[1]+" "+matches[2])
	}
	return results
}

func BuildInitialState(in io.Reader) State {
	s := &StateV1{}
	helpers.ScanLines(in, func(line string) {
		if strings.Contains(line, "nothing relevant") {
			return
		}
		floorNo := readFloor(line)
		s.floors[floorNo] = readItems(line)
	})
	s.setHash()
	return s
}

func FewestMoves(initial State) (int, bool) {
	if initial.Complete() {
		return 0, true
	}

	var (
		seenStates  = map[uint32]bool{initial.Hash(): true}
		previousSet = []State{initial}
		currentSet  = []State{}
		totalMoves  = 0
	)
	for depth := 1; true; depth++ {
		for _, state := range previousSet {
			for newState := range state.AvailableMoves() {
				_, seen := seenStates[newState.Hash()]
				if !seen {
					currentSet = append(currentSet, newState)
					seenStates[newState.Hash()] = true
				}
			}
		}
		totalMoves += len(currentSet)
		fmt.Printf("Depth: %4d, available moves: %4d, total available moves: %6d\n", depth, len(currentSet), totalMoves)
		if len(currentSet) == 0 {
			break
		}
		for _, state := range currentSet {
			if state.Complete() {
				return depth, true
			}
			seenStates[state.Hash()] = true
		}
		previousSet = currentSet
		currentSet = nil
	}
	return 0, false
}

func main() {
	s := BuildInitialState(os.Stdin)
	moves, complete := FewestMoves(s)
	if complete {
		fmt.Println("Min moves:", moves)
	} else {
		fmt.Println("Failed to find solution")
	}
}
