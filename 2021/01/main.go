package main

import (
	"fmt"
	"math"
	"os"

	"github.com/alext/aoc/helpers"
)

func sumInts(in []int) int {
	sum := 0
	for _, i := range in {
		sum += i
	}
	return sum
}

func main() {
	numIncreases := 0
	numWindowIncreases := 0

	window := make([]int, 0)
	previous := math.MaxInt
	helpers.ScanLines(os.Stdin, func(line string) {
		i := helpers.MustAtoi(line)
		if i > previous {
			numIncreases++
		}
		previous = i

		window = append(window, i)
		if len(window) > 3 {
			if sumInts(window[1:]) > sumInts(window[:3]) {
				numWindowIncreases++
			}
			window = window[1:]
		}
	})

	fmt.Println("Num Increases", numIncreases)
	fmt.Println("Num Window Increases", numWindowIncreases)
}
