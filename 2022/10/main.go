package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Instruction struct {
	Op  string
	Arg int
}

func ParseInstruction(input string) Instruction {
	op, arg, _ := strings.Cut(input, " ")
	switch op {
	case "addx":
		return Instruction{Op: op, Arg: helpers.MustAtoi(arg)}
	case "noop":
		return Instruction{Op: op}
	default:
		panic("Failed to parse instruction: " + input)
	}
}

func (i Instruction) Execute(c *CPU) {
	if i.Op == "addx" {
		c.X += i.Arg
	}
}

type CPU struct {
	Instructions []Instruction
	X            int
	Cycles       int
	screen       []bool
}

func NewCPU() *CPU {
	return &CPU{X: 1}
}

func (c *CPU) AddInstruction(i Instruction) {
	if i.Op == "addx" {
		// Add an extra no-op instruction to simulate the 2-cycle nature of add
		c.Instructions = append(c.Instructions, Instruction{Op: "noop-addx"})
	}
	c.Instructions = append(c.Instructions, i)
}

func (c *CPU) ExecuteTo(target int) {
	fmt.Printf("ExecuteTo(%d). Cycles: %d\n", target, c.Cycles)
	if target < c.Cycles {
		panic("Can't go backwards")
	}
	if target > len(c.Instructions) {
		panic("Not enough instructions")
	}
	for c.Cycles < target {
		screenPos := c.Cycles // Screen pos is 0 indexed
		if helpers.AbsInt(c.X-(screenPos%40)) <= 1 {
			c.screen = append(c.screen, true)
		} else {
			c.screen = append(c.screen, false)
		}
		c.Cycles++
		i := c.Instructions[c.Cycles-1]
		fmt.Printf("Executing %v (cycle %d)\n", i, c.Cycles)
		i.Execute(c)
	}
}

func (c *CPU) SignalStrengthDuringCycle(cycle int) int {
	c.ExecuteTo(cycle - 1)
	strength := c.X * cycle
	c.ExecuteTo(cycle)
	return strength
}

func (c *CPU) ScreenOutput() string {
	s := &strings.Builder{}
	for i, lit := range c.screen {
		if i > 0 && i%40 == 0 {
			s.WriteString("\n")
		}
		if lit {
			s.WriteRune('#')
		} else {
			s.WriteRune('.')
		}
	}
	return s.String()
}

func main() {
	c := NewCPU()

	helpers.ScanLines(os.Stdin, func(line string) {
		c.AddInstruction(ParseInstruction(line))
	})

	total := 0
	for n := 20; n <= 220; n += 40 {
		total += c.SignalStrengthDuringCycle(n)
	}
	fmt.Println("Total of signal strengths", total)

	c.ExecuteTo(240)
	fmt.Println(c.ScreenOutput())
}
