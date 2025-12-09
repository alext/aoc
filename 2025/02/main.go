package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

var (
	repeatedDigits      = pcre.MustCompile(`^(\d+)\1$`, 0)
	repeatedDigitsMulti = pcre.MustCompile(`^(\d+)\1+$`, 0)
)

type Range struct {
	Start int
	End   int
}

func (r Range) FindInvalid() []int {
	var invalid []int
	for i := r.Start; i <= r.End; i++ {
		if !repeatedDigits.MatcherString(strconv.Itoa(i), 0).Matches() {
			continue
		}
		invalid = append(invalid, i)

	}
	return invalid
}

func (r Range) FindInvalid2() []int {
	var invalid []int
	for i := r.Start; i <= r.End; i++ {
		if !repeatedDigitsMulti.MatcherString(strconv.Itoa(i), 0).Matches() {
			continue
		}
		invalid = append(invalid, i)

	}
	return invalid
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var ranges []Range
	for _, rawRange := range helpers.SplitCSV(string(input)) {
		start, end, ok := strings.Cut(rawRange, "-")
		if !ok {
			log.Fatalln("Malformed range", rawRange)
		}
		ranges = append(ranges, Range{
			Start: helpers.MustAtoi(start),
			End:   helpers.MustAtoi(end),
		})
	}

	invalidTotal := 0
	invalid2Total := 0
	for _, r := range ranges {
		for _, invalid := range r.FindInvalid() {
			invalidTotal += invalid
		}
		for _, invalid := range r.FindInvalid2() {
			invalid2Total += invalid
		}
	}
	fmt.Println("Invalid total", invalidTotal)
	fmt.Println("Invalid multi total", invalid2Total)
}
