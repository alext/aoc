package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/alext/aoc/helpers"
)

const (
	clothSize = 1000
	conflict  = -1
)

type Cloth struct {
	squares [clothSize][clothSize]int
	claims  map[int]*Claim
}

func NewCloth() *Cloth {
	return &Cloth{
		claims: make(map[int]*Claim),
	}
}

func (cl *Cloth) Reserve(c *Claim) {
	cl.claims[c.ID] = c
	for i := c.Left; i < c.Left+c.Width; i++ {
		for j := c.Top; j < c.Top+c.Height; j++ {
			switch cl.squares[i][j] {
			case 0:
				cl.squares[i][j] = c.ID
			case conflict:
				c.Conflict = true
			default:
				cl.claims[cl.squares[i][j]].Conflict = true
				c.Conflict = true
				cl.squares[i][j] = conflict
			}
		}
	}
}

func (cl *Cloth) Overlaps() uint {
	var dups uint
	for i := 0; i < len(cl.squares); i++ {
		for j := 0; j < len(cl.squares[i]); j++ {
			if cl.squares[i][j] == conflict {
				dups++
			}
		}
	}
	return dups
}

func (cl *Cloth) NonConflicting() []*Claim {
	var results []*Claim
	for _, c := range cl.claims {
		if !c.Conflict {
			results = append(results, c)
		}
	}
	return results
}

type Claim struct {
	ID       int
	Left     int
	Top      int
	Width    int
	Height   int
	Conflict bool
}

// #123 @ 3,2: 5x4
var claimRE = regexp.MustCompile(`#(\d+)\s+@\s+(\d+),(\d+):\s+(\d+)x(\d+)`)

func ParseClaim(input string) *Claim {
	matches := claimRE.FindStringSubmatch(input)
	if matches == nil || len(matches) != 6 {
		log.Fatal("Failed to parse claim", input)
	}
	return &Claim{
		ID:     helpers.MustAtoi(matches[1]),
		Left:   helpers.MustAtoi(matches[2]),
		Top:    helpers.MustAtoi(matches[3]),
		Width:  helpers.MustAtoi(matches[4]),
		Height: helpers.MustAtoi(matches[5]),
	}
}

func main() {
	var cloth = NewCloth()
	helpers.ScanLines(os.Stdin, func(line string) {
		c := ParseClaim(line)
		cloth.Reserve(c)
	})
	fmt.Println("Overlapping squares:", cloth.Overlaps())
	for _, c := range cloth.NonConflicting() {
		fmt.Printf("Claim %d does not conflict\n", c.ID)
	}
}
