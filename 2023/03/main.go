package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alext/aoc/helpers"
)

type Item struct {
	Number int
	Symbol string
	X, Y   int
}

func (i Item) String() string {
	if i.Number > 0 {
		return strconv.Itoa(i.Number)
	}
	return i.Symbol
}

type Schematic struct {
	Positions [][]*Item
	Symbols   []*Item
}

func main() {
	s := &Schematic{}
	helpers.ScanLines(os.Stdin, func(line string) {
		lineIndex := len(s.Positions)
		linePositions := make([]*Item, len(line))

		var lastNum *Item

		for i, ch := range line {
			switch ch {
			case '.':
				lastNum = nil
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if lastNum == nil {
					lastNum = &Item{}
				}
				lastNum.Number *= 10
				lastNum.Number += int(ch - '0')
				linePositions[i] = lastNum
			default:
				lastNum = nil

				sym := &Item{Symbol: string(ch), X: i, Y: lineIndex}
				linePositions[i] = sym
				s.Symbols = append(s.Symbols, sym)
			}
		}

		s.Positions = append(s.Positions, linePositions)
	})

	symbolNumbers := make(map[*Item]bool)
	for _, sym := range s.Symbols {
		for y := helpers.Max(sym.Y-1, 0); y <= helpers.Min(sym.Y+1, len(s.Positions)-1); y++ {
			for x := helpers.Max(sym.X-1, 0); x <= helpers.Min(sym.X+1, len(s.Positions[y])-1); x++ {
				item := s.Positions[y][x]
				if item != nil && item.Number > 0 {
					symbolNumbers[item] = true
				}
			}
		}
	}
	total := 0
	for item := range symbolNumbers {
		total += item.Number
	}
	fmt.Println(total)
}
