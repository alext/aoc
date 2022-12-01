package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

func countDigits(in []string) ([]int, []int) {
	zeroCounts := make([]int, len(in[0]))
	oneCounts := make([]int, len(in[0]))

	for _, line := range in {
		for i, digit := range line {
			switch digit {
			case '0':
				zeroCounts[i]++
			case '1':
				oneCounts[i]++
			default:
				log.Fatalf("Unexpected digit %s in line %s", digit, line)
			}
		}
	}
	return zeroCounts, oneCounts
}

func countDigitsIndex(in []string, index int) (int, int) {
	zeroCount := 0
	oneCount := 0

	for _, line := range in {
		switch line[index] {
		case '0':
			zeroCount++
		case '1':
			oneCount++
		default:
			log.Fatalf("Unexpected digit %s in line %s", line[index], line)
		}
	}
	return zeroCount, oneCount
}

func filterList(list []string, criteria func(string) bool) []string {
	result := make([]string, 0)
	for _, line := range list {
		if criteria(line) {
			result = append(result, line)
		}
	}
	return result
}

func findRating(lines []string, criteria func(int, int, byte) bool) string {

	for i := 0; i < len(lines[0]); i++ {
		zeroCount, oneCount := countDigitsIndex(lines, i)
		lines = filterList(lines, func(line string) bool {
			return criteria(zeroCount, oneCount, line[i])
		})
		if len(lines) == 1 {
			return lines[0]
		}
	}
	log.Fatalf("Failed to filter list. result: %v", lines)
	return "" // unreachable
}

func bin2int(in string) int {
	n, err := strconv.ParseInt(in, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}

func main() {

	lines := make([]string, 0)
	helpers.ScanLines(os.Stdin, func(line string) { lines = append(lines, line) })

	zeroCounts, oneCounts := countDigits(lines)

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

	gamma := bin2int(gammaStr.String())
	epsilon := bin2int(gammaStr.String())

	fmt.Printf("Gamma: %d, Epsilon: %d, product: %d\n", gamma, epsilon, gamma*epsilon)

	oxygenRating := bin2int(findRating(lines, func(zeroCount, oneCount int, digit byte) bool {
		if oneCount >= zeroCount {
			return digit == '1'
		} else {
			return digit == '0'
		}
	}))
	co2Rating := bin2int(findRating(lines, func(zeroCount, oneCount int, digit byte) bool {
		if zeroCount <= oneCount {
			return digit == '0'
		} else {
			return digit == '1'
		}
	}))

	fmt.Printf("O2 rating: %d, CO2 rating: %d, product: %d\n", oxygenRating, co2Rating, oxygenRating*co2Rating)
}
