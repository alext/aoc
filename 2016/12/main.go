package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Instruction struct {
	OpCode string
	Args   []string
}

func (i *Instruction) String() string {
	return i.OpCode + " " + strings.Join(i.Args, " ")
}

func (i *Instruction) Execute(cpu *CPU) {
	switch i.OpCode {
	case "cpy":
		literal, err := strconv.Atoi(i.Args[0])
		if err == nil {
			// integer literal
			cpu.Registers[i.Args[1]] = literal
		} else {
			// register reference
			cpu.Registers[i.Args[1]] = cpu.Registers[i.Args[0]]
		}
	case "inc":
		cpu.Registers[i.Args[0]] += 1
	case "dec":
		cpu.Registers[i.Args[0]] -= 1
	case "jnz":
		value, err := strconv.Atoi(i.Args[0])
		if err != nil {
			// register reference
			value = cpu.Registers[i.Args[0]]
		}
		if value != 0 {
			jump := helpers.MustAtoi(i.Args[1])
			cpu.PC += jump
			return // Prevent normal increment
		}
	default:
		panic("Unrecognised OpCode " + i.OpCode)
	}
	cpu.PC += 1
}

type CPU struct {
	PC           int
	Registers    map[string]int
	Instructions []*Instruction
	Debug        bool
}

func NewCPU() *CPU {
	return &CPU{
		Registers: map[string]int{"a": 0, "b": 0, "c": 0, "d": 0},
	}
}

func (cpu *CPU) String() string {
	return fmt.Sprintf("PC:%2d a:%d b:%d c:%d d:%d", cpu.PC, cpu.Registers["a"], cpu.Registers["b"], cpu.Registers["c"], cpu.Registers["d"])
}

func (cpu *CPU) LoadProgram(in io.Reader) error {
	helpers.ScanLines(in, func(line string) {
		tokens := strings.Split(line, " ")
		i := &Instruction{
			OpCode: tokens[0],
			Args:   tokens[1:],
		}
		cpu.Instructions = append(cpu.Instructions, i)
	})
	return nil
}

func (cpu *CPU) Run() {
	for cpu.PC < len(cpu.Instructions) {
		if cpu.Debug {
			fmt.Printf("%-9s: %s", cpu.Instructions[cpu.PC], cpu)
		}

		cpu.Instructions[cpu.PC].Execute(cpu)

		if cpu.Debug {
			fmt.Printf(" => %s\n", cpu)
		}
	}
}

func main() {
	cpu := NewCPU()
	err := cpu.LoadProgram(os.Stdin)
	if err != nil {
		log.Fatalln("Error loading program:", err)
	}
	cpu.Registers["c"] = 1
	cpu.Run()

	fmt.Println("Final state:", cpu)
}
