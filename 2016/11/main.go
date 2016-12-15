package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"regexp"
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

var itemRE = regexp.MustCompile(`a (.*?)(?:-compatible)? (generator|microchip)`)

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

func (s *State) setHash() {
	s.Hash = 0 // Set to 0 to ensure deterministic hash.
	s.Hash = crc32.ChecksumIEEE([]byte(fmt.Sprintf("%#v", s)))
}

func (s *State) enumerateMoves(ch chan *State) {
	// TODO: Calculate moves

	close(ch)
}

func (s *State) availableMoves() <-chan *State {
	ch := make(chan *State, 2)
	go s.enumerateMoves(ch)
	return ch
}

func (s *State) CountMoves(path []uint32) int {
	if s.Complete() {
		return 0
	}

	path = append(path, s.Hash)

	best := -1000
	for nextState := range s.availableMoves() {
		if alreadyVisited(path, nextState) {
			continue
		}
		candidate := nextState.CountMoves(path)
		if best < 0 || (candidate >= 0 && candidate < best) {
			best = candidate
		}
	}
	return best + 1
}

func alreadyVisited(path []uint32, s *State) bool {
	for _, hash := range path {
		if s.Hash == hash {
			return true
		}
	}
	return false
}

func main() {
	s := BuildInitialState(os.Stdin)
	fmt.Println("Min moves:", s.CountMoves(nil))
}
