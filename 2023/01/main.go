package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/alext/aoc/helpers"
)

func main() {

	total := 0

	firstRe := regexp.MustCompile(`^[^\d]*(\d)`)
	lastRe := regexp.MustCompile(`(\d)[^\d]*$`)

	helpers.ScanLines(os.Stdin, func(line string) {
		matches := firstRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatal("Failed to match first digit in line", line)
		}
		first := matches[1]

		matches = lastRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatal("Failed to match last digit in line", line)
		}

		number := helpers.MustAtoi(first + matches[1])
		total += number
	})

	fmt.Println("Total", total)
}
