package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Valve struct {
	Label      string
	FlowRate   int
	Neighbours []*Valve
	distances  map[string]int
}

func (v *Valve) String() string {
	neighbours := make([]string, 0, len(v.Neighbours))
	for _, n := range v.Neighbours {
		neighbours = append(neighbours, n.Label)
	}
	return fmt.Sprintf("%s flow_rate=%d, neighbours: %s", v.Label, v.FlowRate, strings.Join(neighbours, ", "))
}

func (v *Valve) DistanceTo(other *Valve) int {
	if v.distances == nil {
		v.distances = make(map[string]int)
	}
	if dist, found := v.distances[other.Label]; found {
		return dist
	}
	currentList := map[*Valve]bool{v: true}
	nextList := make(map[*Valve]bool)
	visited := make(map[*Valve]bool)

	for i := 0; len(currentList) > 0; i++ {
		for candidate := range currentList {
			if candidate == other {
				v.distances[other.Label] = i
				return i
			}
			visited[candidate] = true

			for _, n := range candidate.Neighbours {
				if !visited[n] {
					nextList[n] = true
				}
			}
		}
		currentList = nextList
		nextList = make(map[*Valve]bool)
	}

	panic("Failed to find route to " + other.Label)
}

type Move struct {
	TimeRemaining    int
	Current          *Valve
	RemainingValves  map[*Valve]bool
	PressureReleased int
}

func (m *Move) String() string {
	return fmt.Sprintf("%s total:%d", m.Current.Label, m.PressureReleased)
}

func CreateInitialMove(valves map[string]*Valve) *Move {
	m := &Move{
		TimeRemaining:   30,
		Current:         valves["AA"],
		RemainingValves: make(map[*Valve]bool),
	}
	for _, valve := range valves {
		if valve.FlowRate > 0 {
			m.RemainingValves[valve] = true
		}
	}
	return m
}

func (m *Move) CreateNext(nextValve *Valve) *Move {
	distance := m.Current.DistanceTo(nextValve)
	next := &Move{
		TimeRemaining:    m.TimeRemaining - distance - 1, // distance minutes +1 more to open valve
		Current:          nextValve,
		RemainingValves:  make(map[*Valve]bool, len(m.RemainingValves)-1),
		PressureReleased: m.PressureReleased,
	}
	for valve, _ := range m.RemainingValves {
		if valve == nextValve {
			continue
		}
		next.RemainingValves[valve] = true
	}
	if next.TimeRemaining >= 0 {
		next.PressureReleased += nextValve.FlowRate * next.TimeRemaining
	}
	return next
}

func (m *Move) NextMoves() []*Move {
	var moves []*Move
	for v, _ := range m.RemainingValves {
		m := m.CreateNext(v)
		if m.TimeRemaining > 0 {
			moves = append(moves, m)
		}
	}
	return moves
}

func MostPressureRelease(valves map[string]*Valve) int {
	candidateMoves := []*Move{CreateInitialMove(valves)}
	var nextMoves []*Move
	bestResult := 0
	for len(candidateMoves) > 0 {
		for _, move := range candidateMoves {
			if move.PressureReleased > bestResult {
				fmt.Println("New best move", move)
				bestResult = move.PressureReleased
			}
			nextMoves = append(nextMoves, move.NextMoves()...)
		}
		candidateMoves = nextMoves
		nextMoves = nil
	}

	return bestResult
}

func parseInput(in io.Reader) map[string]*Valve {
	valves := make(map[string]*Valve)

	lineRe := regexp.MustCompile(`Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z]+(?:, [A-Z]+)*)`)
	helpers.ScanLines(in, func(line string) {
		matches := lineRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalln("Failed to parse line:", line)
		}
		v := &Valve{
			Label:    matches[1],
			FlowRate: helpers.MustAtoi(matches[2]),
		}
		for _, n := range helpers.SplitCSV(matches[3]) {
			neighbour, ok := valves[n]
			if !ok {
				// Ignore neighbours later in input. Links are bidirectional,
				// so this will be picked up by the reverse link.
				continue
			}
			v.Neighbours = append(v.Neighbours, neighbour)
			neighbour.Neighbours = append(neighbour.Neighbours, v)
		}
		valves[v.Label] = v
	})
	return valves
}

func main() {
	valves := parseInput(os.Stdin)
	fmt.Println(valves)

	fmt.Println("Most pressure release possible:", MostPressureRelease(valves))
}
