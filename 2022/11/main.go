package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Monkey struct {
	Items       []int
	Op          string
	Arg         int
	Divisible   int
	TrueMonkey  int
	FalseMonkey int
	Inspections int
}

func (m Monkey) String() string {
	return fmt.Sprintf(
		"Items: %v\nOp: %s %d\nDivisible %d ? %d : %d\nInspections: %d\n",
		m.Items,
		m.Op, m.Arg,
		m.Divisible, m.TrueMonkey, m.FalseMonkey,
		m.Inspections,
	)
}

func (m *Monkey) TakeTurn(monkeys Monkeys) {
	for _, item := range m.Items {
		worry := 0
		switch m.Op {
		case "+":
			worry = item + m.Arg
		case "*":
			worry = item * m.Arg
		case "sq":
			worry = item * item
		default:
			log.Fatalln("Unexpected Op:", m.Op)
		}
		worry = worry / 3
		if worry%m.Divisible == 0 {
			monkeys[m.TrueMonkey].Items = append(monkeys[m.TrueMonkey].Items, worry)
		} else {
			monkeys[m.FalseMonkey].Items = append(monkeys[m.FalseMonkey].Items, worry)
		}
		m.Inspections++
	}
	m.Items = nil
}

type Monkeys []*Monkey

func (m Monkeys) Round() {
	for i := range m {
		m[i].TakeTurn(m)
	}
}

func main() {
	var monkeys Monkeys
	m := &Monkey{}
	helpers.ScanLines(os.Stdin, func(line string) {
		line = strings.TrimSpace(line)
		if line == "" {
			monkeys = append(monkeys, m)
			m = &Monkey{}
			return
		}
		if line[:6] == "Monkey" {
			return
		}

		property, value, _ := strings.Cut(line, ": ")
		tokens := strings.Split(value, " ")
		switch property {
		case "Starting items":
			for _, num := range helpers.SplitCSV(value) {
				m.Items = append(m.Items, helpers.MustAtoi(num))
			}
		case "Operation":
			if tokens[4] == "old" {
				m.Op = "sq"
			} else {
				m.Op = tokens[3]
				m.Arg = helpers.MustAtoi(tokens[4])
			}
		case "Test":
			m.Divisible = helpers.MustAtoi(tokens[2])
		case "If true":
			m.TrueMonkey = helpers.MustAtoi(tokens[3])
		case "If false":
			m.FalseMonkey = helpers.MustAtoi(tokens[3])
		default:
			log.Fatalln("Unexpected operation:", property, value)
		}
	})
	monkeys = append(monkeys, m)

	fmt.Println("Initial monkeys:", monkeys)

	for i := 0; i < 20; i++ {
		monkeys.Round()
	}

	var inspections []int
	for _, m := range monkeys {
		inspections = append(inspections, m.Inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Println("Inspections:", inspections)
	fmt.Println("Monkey business:", inspections[0]*inspections[1])
}
