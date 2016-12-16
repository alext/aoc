package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/alext/aoc/helpers"
)

const floors = 4

type State struct {
	Floors       [floors][]string
	CurrentFloor int
	Hash         uint32
}

var (
	floorRE = regexp.MustCompile(`The (\S+) floor contains`)
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

var itemRE = regexp.MustCompile(`an? (.*?)(?:-compatible)? (generator|microchip)`)

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

func BuildInitialState(in io.Reader) *State {
	s := &State{}
	helpers.ScanLines(in, func(line string) {
		if strings.Contains(line, "nothing relevant") {
			return
		}
		floorNo := readFloor(line)
		s.Floors[floorNo] = readItems(line)
	})
	s.setHash()
	return s
}

func (s *State) Complete() bool {
	// False if anything on a floor other than the top floor
	for i := 0; i < floors-1; i++ {
		if len(s.Floors[i]) > 0 {
			return false
		}
	}
	// False if not currently on the top floor
	return s.CurrentFloor == floors-1
}

var hashBuffer bytes.Buffer

func (s *State) setHash() {
	// Normalise data for consistency
	for i := 0; i < floors; i++ {
		if s.Floors[i] == nil {
			s.Floors[i] = []string{}
		}
		sort.Strings(s.Floors[i])
	}
	hashBuffer.Reset()
	hashBuffer.WriteByte(byte(s.CurrentFloor))
	for i, f := range s.Floors {
		hashBuffer.WriteByte(byte(i))
		for _, item := range f {
			hashBuffer.Write([]byte(item))
		}
	}
	s.Hash = crc32.ChecksumIEEE(hashBuffer.Bytes())
}

var (
	generators = make([]string, 0)
	microchips = make([]string, 0)
)

func (s *State) Safe() bool {
	for i := 0; i < floors; i++ {
		if len(s.Floors[i]) == 0 {
			continue
		}
		generators = generators[:0]
		microchips = microchips[:0]
		for _, item := range s.Floors[i] {
			if strings.Contains(item, "generator") {
				generators = append(generators, strings.TrimSuffix(item, " generator"))
			} else {
				microchips = append(microchips, strings.TrimSuffix(item, "-compatible microchip"))
			}
		}
		if len(generators) == 0 {
			continue
		}
		for _, chip := range microchips {
			found := false
			for _, generator := range generators {
				if generator == chip {
					found = true
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

func (s *State) Move(newFloor int, items []string) *State {
	move := &State{CurrentFloor: newFloor}
	for i := 0; i < floors; i++ {
		if i == s.CurrentFloor {
			for _, item := range s.Floors[s.CurrentFloor] {
				found := false
				for _, movingItem := range items {
					if item == movingItem {
						found = true
						break
					}
				}
				if !found {
					move.Floors[i] = append(move.Floors[i], item)
				}
			}
			continue
		}
		move.Floors[i] = make([]string, len(s.Floors[i]))
		copy(move.Floors[i], s.Floors[i])
		if i == newFloor {
			move.Floors[i] = append(move.Floors[i], items...)
		}
	}
	move.setHash()
	return move
}

func (s *State) enumerateMoves(ch chan *State) {
	for i, item := range s.Floors[s.CurrentFloor] {
		// Moving a single item
		if s.CurrentFloor < floors-1 {
			candidate := s.Move(s.CurrentFloor+1, []string{item})
			if candidate.Safe() {
				ch <- candidate
			}
		}
		if s.CurrentFloor > 0 {
			candidate := s.Move(s.CurrentFloor-1, []string{item})
			if candidate.Safe() {
				ch <- candidate
			}
		}

		// Moving 2 items
		for _, secondItem := range s.Floors[s.CurrentFloor][i+1:] {
			if s.CurrentFloor < floors-1 {
				candidate := s.Move(s.CurrentFloor+1, []string{item, secondItem})
				if candidate.Safe() {
					ch <- candidate
				}
			}
			if s.CurrentFloor > 0 {
				candidate := s.Move(s.CurrentFloor-1, []string{item, secondItem})
				if candidate.Safe() {
					ch <- candidate
				}
			}
		}
	}

	close(ch)
}

func (s *State) availableMoves() <-chan *State {
	ch := make(chan *State, 2)
	go s.enumerateMoves(ch)
	return ch
}

func FewestMoves(initial *State) (int, bool) {
	if initial.Complete() {
		return 0, true
	}

	var (
		seenStates  = map[uint32]bool{initial.Hash: true}
		previousSet = []*State{initial}
		currentSet  = []*State{}
		totalMoves  = 0
	)
	for depth := 1; true; depth++ {
		for _, state := range previousSet {
			for newState := range state.availableMoves() {
				_, seen := seenStates[newState.Hash]
				if !seen {
					currentSet = append(currentSet, newState)
					seenStates[newState.Hash] = true
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
			seenStates[state.Hash] = true
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
