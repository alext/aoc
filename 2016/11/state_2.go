package main

import (
	"bytes"
	"hash/crc32"
	"sort"
)

type itemPositions [][2]uint8

func (i itemPositions) Len() int      { return len(i) }
func (i itemPositions) Swap(a, b int) { i[a], i[b] = i[b], i[a] }
func (i itemPositions) Less(a, b int) bool {
	if i[a][0] == i[b][0] {
		return i[a][1] < i[b][1]
	}
	return i[a][0] < i[b][0]
}

type StateV2 struct {
	positions    itemPositions
	currentFloor uint8
	hash         uint32
}

func (s *StateV2) Hash() uint32 {
	return s.hash
}

var hashBuffer2 bytes.Buffer

func (s *StateV2) setHash() {
	sort.Sort(s.positions)
	hashBuffer2.Reset()
	hashBuffer2.WriteByte(s.currentFloor)
	for _, pair := range s.positions {
		hashBuffer2.Write(pair[:])
	}
	s.hash = crc32.ChecksumIEEE(hashBuffer2.Bytes())
}

func (s *StateV2) Complete() bool {
	var maxFloor uint8 = floors - 1
	for _, item := range s.positions {
		if item[0] != maxFloor || item[1] != maxFloor {
			return false
		}
	}
	return s.currentFloor == maxFloor
}

func (s *StateV2) Safe() bool {
	generatorFloors := make(map[uint8]bool)
	for _, pair := range s.positions {
		generatorFloors[pair[1]] = true
	}
	for _, pair := range s.positions {
		if pair[0] == pair[1] {
			continue
		}
		if _, ok := generatorFloors[pair[0]]; ok {
			return false
		}
	}
	return true
}

func (s *StateV2) clone() *StateV2 {
	newState := &StateV2{currentFloor: s.currentFloor}
	newState.positions = make(itemPositions, len(s.positions))
	copy(newState.positions, s.positions)
	return newState
}

func (s *StateV2) AvailableMoves() <-chan State {
	ch := make(chan State, 2)

	go func() {
		for i, pair := range s.positions {
			for it, _ := range pair {
				if pair[it] != s.currentFloor {
					continue
				}
				// Moving a single item
				if s.currentFloor < floors-1 {
					candidate := s.clone()
					candidate.positions[i][it] = pair[it] + 1
					candidate.currentFloor = s.currentFloor + 1
					if candidate.Safe() {
						candidate.setHash()
						ch <- candidate
					}
				}

				if s.currentFloor > 0 {
					candidate := s.clone()
					candidate.positions[i][it] = pair[it] - 1
					candidate.currentFloor = s.currentFloor - 1
					if candidate.Safe() {
						candidate.setHash()
						ch <- candidate
					}
				}

				// Moving 2 items
				for j, secondPair := range s.positions[i:] {
					j = j + i
					for jt, _ := range secondPair {
						if i == j && (it == jt || jt == 0) {
							continue
						}
						if secondPair[jt] != s.currentFloor {
							continue
						}
						if s.currentFloor < floors-1 {
							candidate := s.clone()
							candidate.positions[i][it] = pair[it] + 1
							candidate.positions[j][jt] = secondPair[jt] + 1
							candidate.currentFloor = s.currentFloor + 1
							if candidate.Safe() {
								candidate.setHash()
								ch <- candidate
							}
						}

						if s.currentFloor > 0 {
							candidate := s.clone()
							candidate.positions[i][it] = pair[it] - 1
							candidate.positions[j][jt] = secondPair[jt] - 1
							candidate.currentFloor = s.currentFloor - 1
							if candidate.Safe() {
								candidate.setHash()
								ch <- candidate
							}
						}
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}
