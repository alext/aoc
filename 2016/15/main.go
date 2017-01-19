package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

type Disc struct {
	Index     int
	Positions int
	Start     int
}

func (d Disc) AdjustedPosition(time int) int {
	return (d.Start + time + d.Index) % d.Positions
}

func discsAligned(discs []*Disc, time int) bool {
	for _, d := range discs {
		if d.AdjustedPosition(time) != 0 {
			return false
		}
	}
	return true
}

func main() {
	var discs []*Disc

	helpers.ScanLines(os.Stdin, func(line string) {
		d := &Disc{}
		_, err := fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &d.Index, &d.Positions, &d.Start)
		if err != nil {
			log.Fatal("Failed to match input line " + line)
		}
		discs = append(discs, d)
	})

	for time := 0; true; time++ {
		if discsAligned(discs, time) {
			fmt.Println("Discs alligned at time:", time)
			break
		}
	}

	discs = append(discs, &Disc{
		Index:     discs[len(discs)-1].Index + 1,
		Positions: 11,
		Start:     0,
	})

	for time := 0; true; time++ {
		if discsAligned(discs, time) {
			fmt.Println("Discs alligned with extra disc at time:", time)
			break
		}
	}
}
