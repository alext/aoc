package main

import (
	"flag"
	"fmt"

	"github.com/alext/aoc/helpers"
)

type Pos struct {
	X int
	Y int
}

func (p Pos) Distance() int {
	return helpers.AbsInt(p.X) + helpers.AbsInt(p.Y)
}

func main() {
	location := flag.Int("location", 0, "The memory location")
	flag.Parse()

	ringWidth, offset := calculateRingSizeAndOffset(*location)
	pos := locationPosition(*location)
	fmt.Println("Ring:", ringWidth, "Offset:", offset)
	fmt.Println("Distance:", pos.Distance())

	data := make(map[Pos]int)
	data[Pos{0, 0}] = 1
	for i := 2; true; i++ {
		p := locationPosition(i)
		data[p] = sumSurrounding(data, p)
		if data[p] > *location {
			fmt.Printf("Value %d at location %d (%d,%d)\n", data[p], i, p.X, p.Y)
			break
		}
		fmt.Println("Value:", data[p])
	}
}

func sumSurrounding(data map[Pos]int, pos Pos) int {
	sum := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			p := Pos{X: pos.X + i, Y: pos.Y + j}
			sum += data[p]
		}
	}
	return sum
}

func locationPosition(location int) (pos Pos) {
	if location == 1 {
		return Pos{0, 0}
	}
	ringSize, offset := calculateRingSizeAndOffset(location)

	ringRadius := (ringSize - 1) / 2

	side := (offset - 1) / (ringSize - 1)
	sidePos := (offset-1)%(ringSize-1) - (ringRadius - 1)

	switch side {
	case 0:
		pos.X = ringRadius
		pos.Y = sidePos
	case 1:
		pos.X = -sidePos
		pos.Y = ringRadius
	case 2:
		pos.X = -ringRadius
		pos.Y = -sidePos
	case 3:
		pos.X = sidePos
		pos.Y = -ringRadius
	default:
		panic("Square with more than 4 sides")
	}
	return pos
}

func calculateRingSizeAndOffset(location int) (width int, remainder int) {
	remainder = location
	for width = 1; true; width += 2 {
		size := ringSizeForWidth(width)
		if size >= remainder {
			break
		}
		remainder -= size
	}
	return width, remainder
}

func ringSizeForWidth(width int) int {
	if width == 1 {
		return 1
	}
	return width*4 - 4
}
