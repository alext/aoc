package main

import (
	"bytes"
	"hash/crc32"
	"sort"
	"strings"
)

type StateV1 struct {
	floors       [floors][]string
	currentFloor int
	hash         uint32
}

func (s *StateV1) Hash() uint32 {
	return s.hash
}

func (s *StateV1) Complete() bool {
	// False if anything on a floor other than the top floor
	for i := 0; i < floors-1; i++ {
		if len(s.floors[i]) > 0 {
			return false
		}
	}
	// False if not currently on the top floor
	return s.currentFloor == floors-1
}

var hashBuffer bytes.Buffer

func (s *StateV1) setHash() {
	// Normalise data for consistency
	for i := 0; i < floors; i++ {
		if s.floors[i] == nil {
			s.floors[i] = []string{}
		}
		sort.Strings(s.floors[i])
	}
	hashBuffer.Reset()
	hashBuffer.WriteByte(byte(s.currentFloor))
	for i, f := range s.floors {
		hashBuffer.WriteByte(byte(i))
		for _, item := range f {
			hashBuffer.Write([]byte(item))
		}
	}
	s.hash = crc32.ChecksumIEEE(hashBuffer.Bytes())
}

var (
	generators = make([]string, 0)
	microchips = make([]string, 0)
)

func (s *StateV1) Safe() bool {
	for i := 0; i < floors; i++ {
		if len(s.floors[i]) == 0 {
			continue
		}
		generators = generators[:0]
		microchips = microchips[:0]
		for _, item := range s.floors[i] {
			if strings.Contains(item, "generator") {
				generators = append(generators, strings.TrimSuffix(item, " generator"))
			} else {
				microchips = append(microchips, strings.TrimSuffix(item, " microchip"))
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

func (s *StateV1) move(newFloor int, items []string) *StateV1 {
	move := &StateV1{currentFloor: newFloor}
	for i := 0; i < floors; i++ {
		if i == s.currentFloor {
			move.floors[i] = make([]string, 0, len(s.floors[i]))
			for _, item := range s.floors[s.currentFloor] {
				found := false
				for _, movingItem := range items {
					if item == movingItem {
						found = true
						break
					}
				}
				if !found {
					move.floors[i] = append(move.floors[i], item)
				}
			}
			continue
		}
		move.floors[i] = make([]string, len(s.floors[i]))
		copy(move.floors[i], s.floors[i])
		if i == newFloor {
			move.floors[i] = append(move.floors[i], items...)
		}
	}
	move.setHash()
	return move
}

func (s *StateV1) enumerateMoves(ch chan State) {
	for i, item := range s.floors[s.currentFloor] {
		// Moving a single item
		if s.currentFloor < floors-1 {
			candidate := s.move(s.currentFloor+1, []string{item})
			if candidate.Safe() {
				ch <- candidate
			}
		}
		if s.currentFloor > 0 {
			candidate := s.move(s.currentFloor-1, []string{item})
			if candidate.Safe() {
				ch <- candidate
			}
		}

		// Moving 2 items
		for _, secondItem := range s.floors[s.currentFloor][i+1:] {
			if s.currentFloor < floors-1 {
				candidate := s.move(s.currentFloor+1, []string{item, secondItem})
				if candidate.Safe() {
					ch <- candidate
				}
			}
			if s.currentFloor > 0 {
				candidate := s.move(s.currentFloor-1, []string{item, secondItem})
				if candidate.Safe() {
					ch <- candidate
				}
			}
		}
	}

	close(ch)
}

func (s *StateV1) AvailableMoves() <-chan State {
	ch := make(chan State, 2)
	go s.enumerateMoves(ch)
	return ch
}
