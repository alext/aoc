package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

func maxJoultage(bank string, count int) int {
	total := 0
	lastPos := -1
	for count > 0 {
		maxValue := 0
		for i := lastPos + 1; i < len(bank)-count+1; i++ {
			value := int(bank[i] - '0')
			if value > maxValue {
				maxValue = value
				lastPos = i
			}
		}
		total = total*10 + maxValue
		count--
	}

	return total
}

func main() {
	total2 := 0
	total12 := 0
	helpers.ScanLines(os.Stdin, func(bank string) {
		total2 += maxJoultage(bank, 2)
		total12 += maxJoultage(bank, 12)
	})
	fmt.Println("Total2", total2)
	fmt.Println("Total12", total12)
}
