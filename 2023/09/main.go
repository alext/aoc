package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

func allZero(numbers []int) bool {
	for _, n := range numbers {
		if n != 0 {
			return false
		}
	}
	return true
}

func nextNumber(input []int) int {
	if allZero(input) {
		return 0
	}
	var intervals []int
	for i := 1; i < len(input); i++ {
		intervals = append(intervals, input[i]-input[i-1])
	}
	nextInterval := nextNumber(intervals)
	return input[len(input)-1] + nextInterval
}

func previousNumber(input []int) int {
	if allZero(input) {
		return 0
	}
	var intervals []int
	for i := 1; i < len(input); i++ {
		intervals = append(intervals, input[i]-input[i-1])
	}
	previousInterval := previousNumber(intervals)
	return input[0] - previousInterval
}

func main() {
	totalNext := 0
	totalPrevious := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		var numbers []int
		for _, ch := range strings.Fields(line) {
			numbers = append(numbers, helpers.MustAtoi(ch))
		}
		totalNext += nextNumber(numbers)
		totalPrevious += previousNumber(numbers)
	})

	fmt.Println("Total next:", totalNext)
	fmt.Println("Total previous:", totalPrevious)
}
