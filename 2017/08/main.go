package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type instruction struct {
	register  string
	increment bool
	amount    int

	testRegister string
	operator     string
	operand      int
}

// b inc 5 if a > 1
func parseInstruction(line string) (*instruction, error) {
	tokens := strings.Split(line, " ")
	if len(tokens) != 7 {
		return nil, fmt.Errorf("Unexpected number of tokens (%d) in line %s", len(tokens), line)
	}
	i := &instruction{
		register:     tokens[0],
		increment:    tokens[1] == "inc",
		amount:       helpers.MustAtoi(tokens[2]),
		testRegister: tokens[4],
		operator:     tokens[5],
		operand:      helpers.MustAtoi(tokens[6]),
	}
	return i, nil
}

func testConditional(operator string, a, b int) bool {
	switch operator {
	case "==":
		return a == b
	case "!=":
		return a != b
	case ">":
		return a > b
	case ">=":
		return a >= b
	case "<":
		return a < b
	case "<=":
		return a <= b
	default:
		panic("unrecognised operator " + operator)
	}
}

func (i *instruction) Execute(registers map[string]int) {
	if testConditional(i.operator, registers[i.testRegister], i.operand) {
		if i.increment {
			registers[i.register] += i.amount
		} else {
			registers[i.register] -= i.amount
		}
	}
}

func maxRegister(registers map[string]int) (string, int) {
	var maxReg string
	var maxValue int = math.MinInt32
	for reg, value := range registers {
		if value > maxValue {
			maxValue = value
			maxReg = reg
		}
	}
	return maxReg, maxValue
}

func main() {
	registers := make(map[string]int)
	var maxProcessingReg string
	var maxProcessingValue int
	helpers.ScanLines(os.Stdin, func(line string) {
		i, err := parseInstruction(line)
		if err != nil {
			log.Fatal(err)
		}
		i.Execute(registers)
		r, v := maxRegister(registers)
		if v > maxProcessingValue {
			maxProcessingReg, maxProcessingValue = r, v
		}
	})
	maxReg, maxValue := maxRegister(registers)
	fmt.Printf("Max value is %d in register %s\n", maxValue, maxReg)
	fmt.Printf("Max processing value is %d in register %s\n", maxProcessingValue, maxProcessingReg)
}
