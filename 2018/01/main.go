package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

func firstRepeat(changes []int) int {
	frequency := 0
	seen := make(map[int]bool)
	seen[0] = true
	for {
		for _, c := range changes {
			frequency += c
			if seen[frequency] {
				return frequency
			}
			seen[frequency] = true
		}
	}
}

func main() {
	frequency := 0
	changes := make([]int, 0)
	helpers.ScanLines(os.Stdin, func(line string) {
		i := helpers.MustAtoi(line)
		changes = append(changes, i)
		frequency += i
	})
	fmt.Println("Frequency:", frequency)

	fmt.Println("First repeat:", firstRepeat(changes))
}
