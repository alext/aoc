package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

func main() {
	checksum := 0
	divisibleTotal := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		row := make([]int, 0, 0)
		min := math.MaxInt32
		max := 0
		helpers.ScanWrapper(strings.NewReader(line), bufio.ScanWords, func(word string) {
			num := helpers.MustAtoi(word)
			row = append(row, num)
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		})
		checksum += max - min

	Loop:
		for i, _ := range row {
			for j, _ := range row {
				if i == j {
					continue
				}
				if row[i]%row[j] == 0 {
					divisibleTotal += row[i] / row[j]
					break Loop
				}
			}
		}
	})
	fmt.Println("Checksum:", checksum)
	fmt.Println("evenly divisible values:", divisibleTotal)
}
