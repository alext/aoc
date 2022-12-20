package main

import (
	"strings"
	"testing"
)

func TestDistanceTo(t *testing.T) {
	const exampleInput = `Valve AA has flow rate=0; tunnels lead to valves BB, EE, II
Valve BB has flow rate=3; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=0; tunnels lead to valves CC, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD, AA
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=2; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=1; tunnel leads to valve II`
	valves := parseInput(strings.NewReader(exampleInput))

	tests := []struct {
		From     string
		To       string
		Expected int
	}{
		{From: "AA", To: "BB", Expected: 1},
		{From: "AA", To: "CC", Expected: 2},
		{From: "AA", To: "DD", Expected: 2},
		{From: "DD", To: "HH", Expected: 4},
		{From: "AA", To: "AA", Expected: 0},
	}
	for _, test := range tests {
		from := valves[test.From]
		to := valves[test.To]
		actual := from.DistanceTo(to)
		if actual != test.Expected {
			t.Errorf("Expected distance:%d, got:%d for %s->%s", test.Expected, actual, test.From, test.To)
		}
	}
}
