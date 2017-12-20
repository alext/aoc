package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/helpers"
)

type Program struct {
	Instructions []int
	PC           int
}

func (p *Program) Step() {
	nextPC := p.PC + p.Instructions[p.PC]
	if p.Instructions[p.PC] >= 3 {
		p.Instructions[p.PC]--
	} else {
		p.Instructions[p.PC]++
	}
	p.PC = nextPC
}

func (p *Program) Escaped() bool {
	return p.PC < 0 || p.PC >= len(p.Instructions)
}

func main() {
	p := &Program{}
	helpers.ScanLines(os.Stdin, func(line string) {
		p.Instructions = append(p.Instructions, helpers.MustAtoi(line))
	})

	for count := 1; true; count++ {
		p.Step()
		if p.Escaped() {
			fmt.Printf("Escaped in %d steps\n", count)
			break
		}
	}
}
