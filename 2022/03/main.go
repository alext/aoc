package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

type Item uint8

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
	items []Item
}

func NewRucksack(input string) *Rucksack {
	if len(input)%2 != 0 {
		log.Fatalln("Not an even number of items", input)
	}
	var r Rucksack
	for i := 0; i < len(input); i++ {
		item := ParseItem(input[i])
		r.items = append(r.items, item)
	}
	return &r
}

func (r Rucksack) CommonItem() Item {
	half := len(r.items) / 2
	for i := 0; i < half; i++ {
		for j := half; j < len(r.items); j++ {
			if r.items[i] == r.items[j] {
				return r.items[i]
			}
		}
	}
	panic("No common items in rucksack")
}

type Group []*Rucksack

func (g Group) CommonItem() Item {
	if len(g) != 3 {
		log.Fatalln("Group should have 3 members, got", len(g))
	}
	for _, i := range g[0].items {
		for _, j := range g[1].items {
			if i == j {
				for _, k := range g[2].items {
					if i == k {
						return k
					}
				}
			}
		}
	}
	panic("No common item in group")
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
