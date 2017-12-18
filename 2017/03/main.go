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
	distance := calculateDistance(ringWidth, offset)
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

	if side%2 == 0 {
		y = sidePos
		x = ringRadius
		if side >= 2 {
			x = -x
		}
	} else {
		x = sidePos
		y = ringRadius
		if side >= 2 {
			y = -y
		}
	}

	return x, y
}

func calculateDistance(ringSize, offset int) int {
	axisRadius := (ringSize - 1) / 2
	return axisRadius + calculateDistanceFromAxis(ringSize, offset)
}

func calculateDistanceFromAxis(ringSize, offset int) int {
	if ringSize == 1 {
		return 0
	}
	ringRadius := (ringSize - 1) / 2
	sidePos := offset % (ringSize - 1)

	return helpers.AbsInt(sidePos - ringRadius)
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
