package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Card struct {
	Number  int
	Count   int
	Winners map[int]bool
	Numbers []int
}

func ParseCard(input string) *Card {
	// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
	card := &Card{
		Count:   1,
		Winners: make(map[int]bool),
	}

	cardNum, rest, _ := strings.Cut(input, `:`)
	_, err := fmt.Sscanf(cardNum, "Card %d", &card.Number)
	if err != nil {
		log.Fatal(err)
	}

	winners, numbers, _ := strings.Cut(rest, `|`)
	for _, num := range strings.Fields(winners) {
		card.Winners[helpers.MustAtoi(num)] = true
	}
	for _, num := range strings.Fields(numbers) {
		card.Numbers = append(card.Numbers, helpers.MustAtoi(num))
	}

	return card
}

func (c *Card) Score() int {
	score := 0
	for _, num := range c.Numbers {
		if c.Winners[num] {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func (c *Card) NumMatches() int {
	count := 0
	for _, num := range c.Numbers {
		if c.Winners[num] {
			count++
		}
	}
	return count
}

func main() {

	cards := make([]*Card, 0)
	total := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		card := ParseCard(line)
		cards = append(cards, card)
		score := card.Score()
		fmt.Printf("Card %d: score %d\n", card.Number, score)
		total += score
	})
	fmt.Println("Total:", total)

	for i, card := range cards {
		for n := card.NumMatches(); n > 0; n-- {
			cards[i+n].Count += card.Count
		}
	}
	totalCount := 0
	for _, card := range cards {
		fmt.Printf("Card %d: count %d\n", card.Number, card.Count)
		totalCount += card.Count
	}
	fmt.Println("TotalCount:", totalCount)
}
