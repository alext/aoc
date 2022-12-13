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

func (m *Monkey) TakeTurn(monkeys Monkeys, productDivisors int, hasRelief bool) {
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
		if hasRelief {
			worry = worry / 3
		}
		// keep worry manageable
		worry = worry % productDivisors
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

func (m Monkeys) Round(productDivisors int, hasRelief bool) {
	for i := range m {
		m[i].TakeTurn(m, productDivisors, hasRelief)
	}
}

func (m Monkeys) MonkeyBusiness() int {
	var inspections []int
	for _, mky := range m {
		inspections = append(inspections, mky.Inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
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
	monkeys2 := make(Monkeys, 0, len(monkeys))
	productDivisors := 1
	for _, m := range monkeys {
		m2 := *m
		monkeys2 = append(monkeys2, &m2)
		productDivisors *= m.Divisible
	}

	fmt.Println("Initial monkeys:", monkeys)

	for i := 0; i < 20; i++ {
		monkeys.Round(productDivisors, true)
	}

	fmt.Println("Monkey business:", monkeys.MonkeyBusiness())

	for i := 1; i <= 10_000; i++ {
		monkeys2.Round(productDivisors, false)
	}

	fmt.Println("Monkey business2:", monkeys2.MonkeyBusiness())
}
