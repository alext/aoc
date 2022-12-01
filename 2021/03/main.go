package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

func main() {

	var (
		zeroCounts []int
		oneCounts  []int
	)

	helpers.ScanLines(os.Stdin, func(line string) {
		digits := strings.Split(line, "")
		if zeroCounts == nil {
			zeroCounts = make([]int, len(digits))
			oneCounts = make([]int, len(digits))
		}
		for i, digit := range digits {
			switch digit {
			case "0":
				zeroCounts[i]++
			case "1":
				oneCounts[i]++
			default:
				log.Fatalf("Unexpected digit %s in line %s", digit, line)
			}
		}
	})

	gammaStr := &strings.Builder{}
	epsilonStr := &strings.Builder{}
	for i, _ := range zeroCounts {
		if oneCounts[i] > zeroCounts[i] {
			gammaStr.WriteString("1")
			epsilonStr.WriteString("0")
		} else {
			gammaStr.WriteString("0")
			epsilonStr.WriteString("1")
		}
	}

	gamma, err := strconv.ParseInt(gammaStr.String(), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(epsilonStr.String(), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gamma: %d, Epsilon: %d, product: %d\n", gamma, epsilon, gamma*epsilon)
}
