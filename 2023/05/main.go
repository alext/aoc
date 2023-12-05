package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type MapRange struct {
	src    int
	dest   int
	length int
}

type Map struct {
	Ranges []MapRange
}

func ParseMap(input <-chan string) *Map {
	m := &Map{}

	header := <-input
	if header == "" {
		return nil
	}

	for line := range input {
		if line == "" {
			break
		}

		r := MapRange{}
		_, err := fmt.Sscanln(line, &r.dest, &r.src, &r.length)
		if err != nil {
			log.Fatal(err)
		}
		m.Ranges = append(m.Ranges, r)
	}
	return m
}

func (m *Map) Lookup(value int) int {
	for _, r := range m.Ranges {
		if value >= r.src && value < r.src+r.length {
			return r.dest + value - r.src
		}
	}
	return value
}

func (m *Map) LookupRange(input Range) []Range {
	results := []Range{}
	inputs := []Range{input}

InputLoop:
	for len(inputs) > 0 {
		input = inputs[0]
		inputs = inputs[1:]

		for _, r := range m.Ranges {
			if (input.Start >= r.src && input.Start < r.src+r.length) || // start inside range
				(input.End >= r.src && input.End < r.src+r.length) || // end inside range
				(input.Start < r.src && input.End >= r.src+r.length) { // input completely overlaps range
				result := Range{
					Start: r.dest + input.Start - r.src,
					End:   r.dest + input.End - r.src,
				}
				undershoot := r.dest - result.Start
				if undershoot > 0 {
					result.Start += undershoot
					inputs = append(inputs, Range{
						Start: input.Start,
						End:   r.src - 1,
					})
					//fmt.Printf("Adding undershoot input %v\n", inputs[len(inputs)-1])
				}
				overshoot := result.End - (r.dest + r.length - 1)
				if overshoot > 0 {
					result.End -= overshoot
					inputs = append(inputs, Range{
						Start: r.src + r.length,
						End:   input.End,
					})
					//fmt.Printf("Adding overshoot input %v\n", inputs[len(inputs)-1])
				}
				//fmt.Printf("Adding result %v\n", result)
				results = append(results, result)
				continue InputLoop
			}
		}
		results = append(results, input)
	}
	return results
}

type Range struct {
	Start int
	End   int
}

func mapsLookup(maps []*Map, value int) int {
	for _, m := range maps {
		value = m.Lookup(value)
	}
	return value
}

func mapsLookupRange(maps []*Map, input Range) []Range {
	current := []Range{input}
	for _, m := range maps {
		next := []Range{}
		for _, r := range current {
			results := m.LookupRange(r)
			next = append(next, results...)
		}
		current = next
	}
	return current
}

func main() {
	inCh := helpers.StreamLines(os.Stdin)

	var seeds []int
	line := <-inCh
	for _, seed := range strings.Fields(line)[1:] {
		seeds = append(seeds, helpers.MustAtoi(seed))
	}
	fmt.Println("Seeds:", seeds)

	if <-inCh != "" {
		log.Fatalln("Missing blank line after seeds")
	}

	var maps []*Map
	for {
		m := ParseMap(inCh)
		if m == nil {
			break
		}
		maps = append(maps, m)
	}

	lowestLocation := 0
	for _, seed := range seeds {
		location := mapsLookup(maps, seed)
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Println("Lowest location:", lowestLocation)

	lowestLocation = 0
	for i := 0; i < len(seeds)-1; i += 2 {
		seedRange := Range{Start: seeds[i], End: seeds[i] + seeds[i+1] - 1}
		fmt.Println("Considering seed range", seedRange)
		locationRanges := mapsLookupRange(maps, seedRange)
		for _, r := range locationRanges {
			if lowestLocation == 0 || r.Start < lowestLocation {
				lowestLocation = r.Start
			}
		}
	}
	fmt.Println("Lowest location 2:", lowestLocation)
}
