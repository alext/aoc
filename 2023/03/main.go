package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alext/aoc/helpers"
)

type Item struct {
	Number          int
	Symbol          string
	X, Y            int
	AdjacentNumbers map[*Item]bool
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
	totalRatios := 0
	for _, sym := range s.Symbols {
		sym.AdjacentNumbers = make(map[*Item]bool)
		for y := helpers.Max(sym.Y-1, 0); y <= helpers.Min(sym.Y+1, len(s.Positions)-1); y++ {
			for x := helpers.Max(sym.X-1, 0); x <= helpers.Min(sym.X+1, len(s.Positions[y])-1); x++ {
				item := s.Positions[y][x]
				if item != nil && item.Number > 0 {
					symbolNumbers[item] = true
					sym.AdjacentNumbers[item] = true
				}
			}
		}
		if sym.Symbol == "*" && len(sym.AdjacentNumbers) == 2 {
			ratio := 1
			for num := range sym.AdjacentNumbers {
				ratio *= num.Number
			}
			totalRatios += ratio
		}
	}
	total := 0
	for item := range symbolNumbers {
		total += item.Number
	}
	fmt.Println("Total:", total)
	fmt.Println("TotalRatios:", totalRatios)
}
