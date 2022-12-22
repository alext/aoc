package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
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
			v.distances[other.Label] = i
			if candidate == other {
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
	ETimeRemaining   int
	ECurrent         *Valve
	RemainingValves  map[*Valve]bool
	PressureReleased int
	Path             string
}

func (m *Move) String() string {
	if m.ECurrent != nil {
		return fmt.Sprintf("%s E:%s total:%d, path:%s", m.Current.Label, m.ECurrent.Label, m.PressureReleased, m.Path)
	} else {
		return fmt.Sprintf("%s total:%d, path:%s", m.Current.Label, m.PressureReleased, m.Path)
	}
}

func CreateInitialMove(valves map[string]*Valve, withElephant bool) *Move {
	m := &Move{
		TimeRemaining:   30,
		Current:         valves["AA"],
		RemainingValves: make(map[*Valve]bool),
		Path:            "AA",
	}
	for _, valve := range valves {
		if valve.FlowRate > 0 {
			m.RemainingValves[valve] = true
		}
	}
	if withElephant {
		m.ECurrent = m.Current
		m.TimeRemaining = 26
		m.ETimeRemaining = 26
	}
	return m
}

func (m *Move) CannotBeat(currentBest int) bool {
	rates := make([]int, 0, len(m.RemainingValves))
	for v, _ := range m.RemainingValves {
		rates = append(rates, v.FlowRate)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rates)))
	potentialBest := m.PressureReleased
	timeRemaining := m.TimeRemaining
	for _, rate := range rates {
		timeRemaining -= 2
		if timeRemaining <= 0 {
			break
		}
		potentialBest += rate * timeRemaining
	}
	return currentBest >= potentialBest
}

func (m *Move) CreateNext(nextValve *Valve, elephantMove bool) *Move {
	next := &Move{
		Current:          m.Current,
		TimeRemaining:    m.TimeRemaining,
		ECurrent:         m.ECurrent,
		ETimeRemaining:   m.ETimeRemaining,
		PressureReleased: m.PressureReleased,
		RemainingValves:  make(map[*Valve]bool, len(m.RemainingValves)-1),
	}
	if elephantMove {
		next.ETimeRemaining = m.ETimeRemaining - m.ECurrent.DistanceTo(nextValve) - 1 // distance minutes +1 more to open valve
		next.ECurrent = nextValve
		next.Path = m.Path + " E:" + nextValve.Label
		if next.ETimeRemaining >= 0 {
			next.PressureReleased += nextValve.FlowRate * next.ETimeRemaining
		}
	} else {
		next.TimeRemaining = m.TimeRemaining - m.Current.DistanceTo(nextValve) - 1 // distance minutes +1 more to open valve
		next.Current = nextValve
		next.Path = m.Path + " " + nextValve.Label
		if next.TimeRemaining >= 0 {
			next.PressureReleased += nextValve.FlowRate * next.TimeRemaining
		}
	}
	for valve, _ := range m.RemainingValves {
		if valve == nextValve {
			continue
		}
		next.RemainingValves[valve] = true
	}
	return next
}

func (m *Move) NextMoves() []*Move {
	var moves []*Move
	for v, _ := range m.RemainingValves {
		if m.TimeRemaining > 0 {
			moves = append(moves, m.CreateNext(v, false))
		}
		if m.ECurrent != nil && m.ETimeRemaining > 0 {
			moves = append(moves, m.CreateNext(v, true))
		}
	}
	return moves
}

func MostPressureRelease(valves map[string]*Valve, withElephant bool) int {
	candidateMoves := []*Move{CreateInitialMove(valves, withElephant)}
	var nextMoves []*Move
	bestResult := 0
	for len(candidateMoves) > 0 {
		for _, move := range candidateMoves {
			//fmt.Println("Considering move", move)
			//fmt.Println(move.RemainingValves)
			if move.PressureReleased > bestResult {
				fmt.Println("New best move", move)
				bestResult = move.PressureReleased
			}
			for _, next := range move.NextMoves() {
				if next.CannotBeat(bestResult) {
					//fmt.Printf("Culling move %s, cannot beat %d\n", next, bestResult)
					continue
				}
				nextMoves = append(nextMoves, next)
			}
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
	//fmt.Println(valves)

	fmt.Println("Most pressure release possible:", MostPressureRelease(valves, false))

	fmt.Println("Most pressure release with elephant:", MostPressureRelease(valves, true))
}
