package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/alext/aoc/helpers"
)

func main() {
	elves := make([]int, 0)

	elfCalories := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		if line == "" {
			elves = append(elves, elfCalories)
			elfCalories = 0
			return
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		elfCalories += n
	})

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Println("Max Calories:", elves[0])
	fmt.Println("Top 3 total:", elves[0]+elves[1]+elves[2])
}
