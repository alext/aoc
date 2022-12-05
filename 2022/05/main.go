package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Ship struct {
	Stacks [][]string
}

func ParseShipContents(lines []string) *Ship {
	if len(lines) < 1 {
		log.Fatalln("Expected at least one stack line")
	}
	s := Ship{}
	s.Stacks = make([][]string, len(strings.Fields(lines[len(lines)-1])))

	stackLineRe := regexp.MustCompile(`.(.). ?`)
	for i := len(lines) - 2; i >= 0; i-- {
		matches := stackLineRe.FindAllStringSubmatch(lines[i], -1)
		if len(matches) != len(s.Stacks) {
			log.Fatalf("Failed to parse line '%s', got '%v'\n", lines[i], matches)
		}
		for j, match := range matches {
			if match[1] != " " {
				s.Stacks[j] = append(s.Stacks[j], match[1])
			}
		}
	}

	return &s
}

func (s *Ship) Clone() *Ship {
	ss := Ship{
		Stacks: make([][]string, len(s.Stacks)),
	}
	for i, stack := range s.Stacks {
		ss.Stacks[i] = append(ss.Stacks[i], stack...)
	}
	return &ss
}

func (s *Ship) Move(quantity, from, to int) {
	// Convert to zero-indexed
	from, to = from-1, to-1
	for quantity > 0 {
		s.Stacks[to] = append(s.Stacks[to], s.Stacks[from][len(s.Stacks[from])-1])
		s.Stacks[from] = s.Stacks[from][:len(s.Stacks[from])-1]
		quantity--
	}
}

func (s *Ship) MoveStacked(quantity, from, to int) {
	// Convert to zero-indexed
	from, to = from-1, to-1

	fromLen := len(s.Stacks[from])
	crates := s.Stacks[from][fromLen-quantity : fromLen]
	s.Stacks[to] = append(s.Stacks[to], crates...)
	s.Stacks[from] = s.Stacks[from][:fromLen-quantity]
}

func (s *Ship) StackTops() string {
	var b strings.Builder
	for _, stack := range s.Stacks {
		if len(stack) > 0 {
			b.WriteString(stack[len(stack)-1])
		}
	}
	return b.String()
}

func main() {
	inputCh := helpers.StreamLines(os.Stdin)

	var stackLines []string
	for line := range inputCh {
		if line == "" {
			break
		}
		stackLines = append(stackLines, line)
	}

	ship := ParseShipContents(stackLines)
	ship2 := ship.Clone()

	for instruction := range inputCh {
		tokens := strings.Fields(instruction)
		quantity := helpers.MustAtoi(tokens[1])
		from := helpers.MustAtoi(tokens[3])
		to := helpers.MustAtoi(tokens[5])
		ship.Move(quantity, from, to)

		ship2.MoveStacked(quantity, from, to)
	}

	fmt.Println("Stack tops:", ship.StackTops())
	fmt.Println("Stack tops 2:", ship2.StackTops())
}
