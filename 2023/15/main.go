package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"

	"github.com/alext/aoc/helpers"
)

func Hash(input string) uint8 {
	value := 0
	for _, r := range input {
		value += int(r)
		value *= 17
		value = value % 256
	}
	return uint8(value)
}

type Lens struct {
	Label       string
	FocalLength int
}

type Box struct {
	Lenses []Lens
}

type Boxes map[uint8][]Lens

func (b Boxes) ProcessInstruction(label, op, arg string) {
	slot := Hash(label)
	if op == "-" {
		b[slot] = slices.DeleteFunc(b[slot], func(l Lens) bool { return l.Label == label })
		return
	}
	// op == "="
	focalLength := helpers.MustAtoi(arg)
	for i := range b[slot] {
		if b[slot][i].Label == label {
			b[slot][i].FocalLength = focalLength
			return
		}
	}
	// Not found, so add.
	b[slot] = append(b[slot], Lens{Label: label, FocalLength: focalLength})
}

func (b Boxes) FocussingPower() int {
	totalPower := 0

	for boxNum, lenses := range b {
		for slot, lens := range lenses {
			totalPower += (int(boxNum) + 1) * (slot + 1) * lens.FocalLength
		}
	}
	return totalPower
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	instructions := helpers.SplitCSV(string(input))
	hashTotal := 0
	for _, inst := range instructions {
		hash := Hash(inst)
		fmt.Printf("Instruction %s, hash:%d\n", inst, hash)
		hashTotal += int(hash)
	}
	fmt.Println("Hash total:", hashTotal)

	boxes := make(Boxes)
	instRe := regexp.MustCompile(`^([a-z]+)([-=])(\d*)$`)
	for _, inst := range instructions {
		matches := instRe.FindStringSubmatch(inst)
		if matches == nil {
			log.Fatalf("Failed to match instruction %s", inst)
		}
		boxes.ProcessInstruction(matches[1], matches[2], matches[3])
	}

	fmt.Println(boxes)
	fmt.Println("Total Focussing power:", boxes.FocussingPower())
}
