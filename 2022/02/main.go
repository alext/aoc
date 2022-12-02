package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Outcome int

const (
	Lose Outcome = 1
	Draw Outcome = 2
	Win  Outcome = 3
)

func ParseOutcome(in string) Outcome {
	switch in {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		panic("Unexpected outcome input " + in)
	}
}

type Play int

const (
	Rock     Play = 1
	Paper    Play = 2
	Scissors Play = 3
)

func ParsePlay(in string) Play {
	switch in {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		panic("Unexpected play input " + in)
	}
}

func (p Play) Beats(other Play) bool {
	switch p {
	case Rock:
		return other == Scissors
	case Paper:
		return other == Rock
	case Scissors:
		return other == Paper
	default:
		panic("unexpected play")
	}
}

func (p Play) ScorePlay(other Play) int {
	score := int(p) // score for shape played
	if p == other {
		// draw
		score += 3
	} else if p.Beats(other) {
		score += 6
	}
	return score
}

func (p Play) PlayAgainstToAchieve(target Outcome) Play {
	if target == Draw {
		return p
	}
	switch p {
	case Rock:
		if target == Win {
			return Paper
		}
		return Scissors
	case Paper:
		if target == Win {
			return Scissors
		}
		return Rock
	case Scissors:
		if target == Win {
			return Rock
		}
		return Paper
	default:
		panic("unexpected play")
	}
}

func main() {
	score := 0
	score2 := 0

	helpers.ScanLines(os.Stdin, func(line string) {
		plays := strings.Split(line, " ")
		if len(plays) != 2 {
			log.Fatal("Expected 2 plays on line", line)
		}

		theirs := ParsePlay(plays[0])
		ours := ParsePlay(plays[1])

		score += ours.ScorePlay(theirs)

		// part 2: second input is the desired outcome
		targetOutcome := ParseOutcome(plays[1])

		ours2 := theirs.PlayAgainstToAchieve(targetOutcome)
		score2 += ours2.ScorePlay(theirs)
	})

	fmt.Println("Final score:", score)
	fmt.Println("Final score2:", score2)
}
