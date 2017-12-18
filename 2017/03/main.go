package main

import (
	"flag"
	"fmt"

	"github.com/alext/aoc/helpers"
)

func main() {
	location := flag.Int("location", 0, "The memory location")
	flag.Parse()

	ringWidth, offset := calculateRingSizeAndOffset(*location)
	x, y := locationPosition(*location)
	distance := helpers.AbsInt(x) + helpers.AbsInt(y)
	fmt.Println("Ring:", ringWidth, "Offset:", offset)
	fmt.Println("Distance:", distance)
}

func locationPosition(location int) (x, y int) {
	if location == 1 {
		return 0, 0
	}
	ringSize, offset := calculateRingSizeAndOffset(location)

	ringRadius := (ringSize - 1) / 2

	side := (offset - 1) / (ringSize - 1)
	sidePos := (offset-1)%(ringSize-1) - (ringRadius - 1)

	switch side {
	case 0:
		x = ringRadius
		y = sidePos
	case 1:
		x = -sidePos
		y = ringRadius
	case 2:
		x = -ringRadius
		y = -sidePos
	case 3:
		x = sidePos
		y = -ringRadius
	default:
		panic("Square with more than 4 sides")
	}
	return x, y
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
