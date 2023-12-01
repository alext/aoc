package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/alext/aoc/helpers"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseNumber(input string) int {
	if len(input) == 1 {
		return helpers.MustAtoi(input)
	}
	value, ok := numbers[input]
	if !ok {
		log.Fatal("Failed to parse input", input)
	}
	return value
}

func main() {

	total := 0

	digits := `\d`
	for word := range numbers {
		digits += `|` + word
	}
	firstRe := regexp.MustCompile(`^.*?(` + digits + `)`)
	lastRe := regexp.MustCompile(`.*(` + digits + `).*?$`)

	helpers.ScanLines(os.Stdin, func(line string) {
		matches := firstRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatal("Failed to match first digit in line", line)
		}
		first := parseNumber(matches[1])

		matches = lastRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatal("Failed to match last digit in line", line)
		}
		last := parseNumber(matches[1])

		number := first*10 + last
		total += number
	})

	fmt.Println("Total", total)
}
