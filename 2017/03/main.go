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
		size := ringSize(width)
		if size >= remainder {
			break
		}
		remainder -= size
	}
	return width, remainder
}

func ringSize(width int) int {
	if width == 1 {
		return 1
	}
	return width*4 - 4
}
