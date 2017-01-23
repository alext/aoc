package main

import (
	"crypto/md5"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"
)

const (
	gridSize      = 4
	openDoorChars = "bcdef"
)

type Position struct {
	X    int
	Y    int
	Path string
}

func (p *Position) IsVault() bool {
	return p.X == gridSize && p.Y == gridSize
}

func (p *Position) AvailableMoves(passcode string) <-chan *Position {
	ch := make(chan *Position, 2)
	go func() {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(passcode+p.Path)))
		if p.Y > 1 && strings.ContainsRune(openDoorChars, rune(hash[0])) {
			ch <- &Position{X: p.X, Y: p.Y - 1, Path: p.Path + "U"}
		}
		if p.Y < gridSize && strings.ContainsRune(openDoorChars, rune(hash[1])) {
			ch <- &Position{X: p.X, Y: p.Y + 1, Path: p.Path + "D"}
		}
		if p.X > 1 && strings.ContainsRune(openDoorChars, rune(hash[2])) {
			ch <- &Position{X: p.X - 1, Y: p.Y, Path: p.Path + "L"}
		}
		if p.X < gridSize && strings.ContainsRune(openDoorChars, rune(hash[3])) {
			ch <- &Position{X: p.X + 1, Y: p.Y, Path: p.Path + "R"}
		}
		close(ch)
	}()
	return ch
}

// Represents an as-yet unknown distance
const Infinity = 999999

func ShortestPath(passcode string) (string, error) {
	start := &Position{X: 1, Y: 1}

	candidates := make(map[*Position]bool)
	candidates[start] = true

	for {
		if len(candidates) == 0 {
			return "", errors.New("No path found")
		}

		var (
			distance  int = Infinity
			candidate *Position
		)
		for pos, _ := range candidates {
			if len(pos.Path) < distance {
				candidate = pos
				distance = len(pos.Path)
			}
		}
		delete(candidates, candidate)
		for pos := range candidate.AvailableMoves(passcode) {
			if pos.IsVault() {
				return pos.Path, nil
			}
			candidates[pos] = true
		}
	}
}

func main() {
	passcode := flag.String("passcode", "", "The passcode")

	flag.Parse()

	path, err := ShortestPath(*passcode)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Shortest path:", path)
}
