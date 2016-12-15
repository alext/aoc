package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

const floors = 4

type Floor []string

type State struct {
	Floors       [floors]Floor
	CurrentFloor int
	Hash         uint32
}

func BuildInitialState(in io.Reader) *State {
	s := &State{}
	// TODO: Parse input
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
	s.Hash = 0
	s.Hash = crc32.ChecksumIEEE([]byte(fmt.Sprintf("%#v", s)))
}

func (s *State) enumerateMoves(ch chan *State) {
	// TODO: Calculate moves

	close(ch)
}

func (s *State) availableMoves() <-chan *State {
	ch := make(chan *State)
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
