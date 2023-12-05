package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Range struct {
	src    int
	dest   int
	length int
}

type Map struct {
	Ranges []Range
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

		r := Range{}
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

func mapsLookup(maps []*Map, value int) int {
	for _, m := range maps {
		value = m.Lookup(value)
	}
	return value
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
		fmt.Println("Looking up seed", seed)
		location := mapsLookup(maps, seed)
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Println("Lowest location:", lowestLocation)
}
