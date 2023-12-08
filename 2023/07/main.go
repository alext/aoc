package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Card int

const Joker Card = 0
const (
	Jack Card = 11 + iota
	Queen
	King
	Ace
)

func ParseCard(c rune) Card {
	switch c {
	case 'T':
		return Card(10)
	case 'J':
		return Joker
	case 'Q':
		return Queen
	case 'K':
		return King
	case 'A':
		return Ace
	default:
		return Card(helpers.MustAtoi(string(c)))
	}
}

func (c Card) String() string {
	switch c {
	case 10:
		return "T"
	case Jack, Joker:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return strconv.Itoa(int(c))
	}
}

//go:generate go run golang.org/x/tools/cmd/stringer -type HandType
type HandType int

const (
	Unknown HandType = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards [5]Card
	Bid   int
	Type  HandType
}

func ParseHand(line string) Hand {
	cards, bid, _ := strings.Cut(line, " ")
	h := Hand{
		Bid: helpers.MustAtoi(bid),
	}
	for i, c := range cards {
		h.Cards[i] = ParseCard(c)
	}
	h.Type = h.calculateType()
	return h
}

func (h Hand) calculateType() HandType {
	counts := make(map[Card]int)
	jokers := 0
	var maxCard Card
	maxCount := 0
	for _, c := range h.Cards {
		if c == Joker {
			jokers++
		} else {
			counts[c]++
			if counts[c] > maxCount {
				maxCard, maxCount = c, counts[c]
			}
		}
	}
	counts[maxCard] += jokers
	countCounts := make(map[int]int)
	for _, count := range counts {
		countCounts[count]++
	}
	switch {
	case countCounts[5] > 0:
		return FiveOfAKind
	case countCounts[4] > 0:
		return FourOfAKind
	case countCounts[3] > 0 && countCounts[2] > 0:
		return FullHouse
	case countCounts[3] > 0:
		return ThreeOfAKind
	case countCounts[2] >= 2:
		return TwoPair
	case countCounts[2] > 0:
		return OnePair
	default:
		return HighCard
	}
}

func cmpCards(a, b [5]Card) int {
	for i := range a {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

func cmpHands(a, b Hand) int {
	if n := cmp.Compare(a.Type, b.Type); n != 0 {
		return n
	}
	return cmpCards(a.Cards, b.Cards)
}

func main() {
	var hands []Hand
	helpers.ScanLines(os.Stdin, func(line string) {
		hands = append(hands, ParseHand(line))
	})

	slices.SortFunc(hands, cmpHands)
	winnings := 0
	for i, h := range hands {
		rank := i + 1
		fmt.Printf("Rank %d: %v\n", rank, h)
		winnings += rank * h.Bid
	}
	fmt.Println(winnings)
}
