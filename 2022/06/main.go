package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

func AllDifferent(in string) bool {
	chars := make(map[rune]bool)
	for _, c := range in {
		chars[c] = true
	}
	return len(chars) == len(in)
}

func main() {
	input := ""

	helpers.ScanLines(os.Stdin, func(line string) {
		if input != "" {
			log.Fatal("Multiple input lines found")
		}
		input = line
	})

	if len(input) < 4 {
		log.Fatalf("Expected at least 4 chars in input. Got %d", len(input))
	}

	i := 4
	for ; i <= len(input); i++ {
		if AllDifferent(input[i-4 : i]) {
			fmt.Printf("Start of packet after %d chars\n", i)
			break
		}
	}

	for i = i + 10; i <= len(input); i++ {
		if AllDifferent(input[i-14 : i]) {
			fmt.Printf("Start of message after %d chars\n", i)
			break
		}
	}
}
