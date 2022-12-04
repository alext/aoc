package main

import (
	"fmt"
	"log"
	"math/bits"
	"os"

	"github.com/alext/aoc/helpers"
)

type Item uint

func ParseItem(letter byte) Item {
	if letter >= 'a' && letter <= 'z' {
		return Item(letter-'a') + 1
	}
	if letter >= 'A' && letter <= 'Z' {
		return Item(letter-'A') + 27
	}
	panic(fmt.Sprintf("Unexpected item letter %s", letter))
}

func (i Item) Priority() int { return int(i) }

func (i Item) String() string {
	if i < 27 {
		return string(i - 1 + 'a')
	} else {
		return string(i - 27 + 'A')
	}
}

type Rucksack struct {
	left  uint
	right uint
}

func NewRucksack(input string) *Rucksack {
	if len(input)%2 != 0 {
		log.Fatalln("Not an even number of items", input)
	}
	var r Rucksack
	half := len(input) / 2
	for i := 0; i < len(input); i++ {
		item := ParseItem(input[i])
		if i < half {
			r.left |= 1 << item
		} else {
			r.right |= 1 << item
		}
	}
	return &r
}

func (r Rucksack) CommonItem() Item {
	combine := r.left & r.right
	if combine == 0 {
		log.Fatal("No common items in rucksack")
	}
	return Item(bits.TrailingZeros(combine))
}

type Group []*Rucksack

func (g Group) CommonItem() Item {
	var combine uint = ^uint(0) // All ones
	for _, r := range g {
		combine &= r.left | r.right
	}
	if combine == 0 {
		log.Fatal("No common items in group")
	}
	return Item(bits.TrailingZeros(combine))
}

func main() {

	score := 0
	groupScore := 0
	group := make(Group, 0, 3)
	helpers.ScanLines(os.Stdin, func(line string) {
		if len(line)%2 != 0 {
			log.Fatalln("Line is not an even number of characters", line)
		}
		r := NewRucksack(line)
		score += r.CommonItem().Priority()
		group = append(group, r)
		if len(group) == 3 {
			groupScore += group.CommonItem().Priority()
			group = group[:0]
		}
	})

	fmt.Println("Score:", score)
	fmt.Println("Group Score:", groupScore)
}
