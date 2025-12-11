package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

func Reduce[S ~[]E, E any, T any](s S, f func(T, E) T, init T) T {
	for _, v := range s {
		init = f(init, v)
	}
	return init
}

func Calculate(op string, numbers []int) int {
	result := 0
	switch op {
	case "+":
		result = Reduce(numbers, func(acc, n int) int { return acc + n }, 0)
	case "*":
		result = Reduce(numbers, func(acc, n int) int { return acc * n }, 1)
	default:
		log.Fatalln("Invalid operator", op)

	}
	return result

}

func parseNumberColumns(lines []string) [][]int {
	var result [][]int

	var group []int
	for c := 0; c < len(lines[0]); c++ {
		value := 0
		for r := 0; r < len(lines); r++ {
			ch := lines[r][c]
			if ch == ' ' {
				continue
			}
			value *= 10
			value += int(ch - '0')
		}

		if value == 0 {
			// end of group
			result = append(result, group)
			group = nil
		} else {
			group = append(group, value)
		}
	}
	result = append(result, group)

	return result
}

func main() {
	var input []string
	helpers.ScanLines(os.Stdin, func(line string) {
		input = append(input, line)
	})

	numberLines := input[:len(input)-1]
	operators := strings.Fields(input[len(input)-1])

	var rowNumbers [][]string
	for _, row := range numberLines {
		rowNumbers = append(rowNumbers, strings.Fields(row))
	}

	numbers := make([]int, len(rowNumbers))
	total := 0
	for c, op := range operators {
		for r, row := range rowNumbers {
			numbers[r] = helpers.MustAtoi(row[c])
		}

		total += Calculate(op, numbers)
	}
	fmt.Println("Total part1:", total)

	colNumbers := parseNumberColumns(numberLines)
	total = 0
	for i, group := range colNumbers {
		total += Calculate(operators[i], group)
	}
	fmt.Println("Total part2:", total)
}
