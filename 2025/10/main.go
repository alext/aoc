package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Machine struct {
	LightsTarget []bool
	Buttons      [][]int
}

type Move struct {
	Lights  []bool
	Presses int
	Button  []int
}

func (m *Move) PressButton() {
	for _, pos := range m.Button {
		m.Lights[pos] = !m.Lights[pos]
	}
	m.Presses++
}

func (m Machine) BestButtonSequence() int {
	var candidates []*Move
	for _, b := range m.Buttons {
		candidates = append(candidates, &Move{
			Lights: make([]bool, len(m.LightsTarget)),
			Button: b,
		})
	}
	for len(candidates) > 0 {
		move := candidates[0]
		candidates = candidates[1:]
		move.PressButton()

		if slices.Equal(move.Lights, m.LightsTarget) {
			return move.Presses
		}

		for _, b := range m.Buttons {
			candidates = append(candidates, &Move{
				Lights:  slices.Clone(move.Lights),
				Presses: move.Presses,
				Button:  b,
			})
		}
	}
	log.Fatalln("Ran out of moves")
	return -1
}

func main() {
	lightRe := regexp.MustCompile(`\[(.+?)\]`)
	buttonRe := regexp.MustCompile(`\((\d+(?:,\d+)*)\)`)

	var machines []Machine
	helpers.ScanLines(os.Stdin, func(line string) {
		match := lightRe.FindStringSubmatch(line)
		m := Machine{
			LightsTarget: make([]bool, len(match[1])),
		}
		for i, ch := range match[1] {
			if ch == '#' {
				m.LightsTarget[i] = true
			}
		}
		matches := buttonRe.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			rawValues := strings.Split(match[1], ",")
			var values []int
			for _, v := range rawValues {
				values = append(values, helpers.MustAtoi(v))
			}
			m.Buttons = append(m.Buttons, values)
		}

		machines = append(machines, m)
	})

	totalPresses := 0
	for _, machine := range machines {
		totalPresses += machine.BestButtonSequence()
	}
	fmt.Println("Total presses:", totalPresses)
}
