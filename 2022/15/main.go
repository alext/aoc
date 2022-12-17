package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Pos = helpers.Pos

type Sensor struct {
	Pos            Pos
	Beacon         Pos
	BeaconDistance int
}

func (s Sensor) String() string {
	return fmt.Sprintf("Sensor: %s, Beacon: %s", s.Pos, s.Beacon)
}

func (s Sensor) PositionClear(p Pos) bool {
	if p == s.Beacon {
		return false
	}
	return s.Pos.DistanceTo(p) <= s.BeaconDistance
}

type Grid struct {
	Sensors []Sensor
	Min     Pos
	Max     Pos
}

func (g *Grid) String() string {
	positions := make(map[Pos]string)
	for _, s := range g.Sensors {
		positions[s.Pos] = "S"
		positions[s.Beacon] = "B"
	}
	b := &strings.Builder{}
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			if ch, ok := positions[Pos{X: x, Y: y}]; ok {
				b.WriteString(ch)
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (g *Grid) expandEdges(s Sensor) {
	if s.Pos.X-s.BeaconDistance < g.Min.X {
		g.Min.X = s.Pos.X - s.BeaconDistance
	}
	if s.Pos.Y-s.BeaconDistance < g.Min.Y {
		g.Min.Y = s.Pos.Y - s.BeaconDistance
	}
	if s.Pos.X+s.BeaconDistance > g.Max.X {
		g.Max.X = s.Pos.X + s.BeaconDistance
	}
	if s.Pos.Y+s.BeaconDistance > g.Max.Y {
		g.Max.Y = s.Pos.Y + s.BeaconDistance
	}
}

func (g *Grid) AddSensor(s Sensor) {
	s.BeaconDistance = s.Pos.DistanceTo(s.Beacon)
	g.Sensors = append(g.Sensors, s)
	g.expandEdges(s)
}

func (g *Grid) CountClearSpots(row int) int {
	count := 0
	for x := g.Min.X; x <= g.Max.X; x++ {
		p := Pos{X: x, Y: row}
		clear := false
		for _, s := range g.Sensors {
			if s.PositionClear(p) {
				clear = true
				break
			}
		}
		if clear {
			count++
		}
	}
	return count
}

func main() {
	checkRow := flag.Int("checkrow", 0, "Row to check for clear spots")
	flag.Parse()

	g := &Grid{}
	helpers.ScanLines(os.Stdin, func(line string) {
		var s Sensor
		_, err := fmt.Sscanf(
			line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.Pos.X, &s.Pos.Y,
			&s.Beacon.X, &s.Beacon.Y,
		)
		if err != nil {
			log.Fatal(err)
		}
		g.AddSensor(s)
	})

	fmt.Println("Min", g.Min)
	fmt.Println("Max", g.Max)

	//fmt.Println(g)
	fmt.Printf("ClearSpots on y=%d: %d\n", *checkRow, g.CountClearSpots(*checkRow))
}
