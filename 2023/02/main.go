package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alext/aoc/helpers"
)

const (
	redMax   = 12
	greenMax = 13
	blueMax  = 14
)

func minCubes(draws []string) map[string]int {
	counts := make(map[string]int, 3)
	for _, draw := range draws {
		sets := strings.Split(draw, `, `)
		for _, set := range sets {
			n, cube, ok := strings.Cut(set, ` `)
			if !ok {
				log.Fatalf("Invalid set %s in draw %s", set, draw)
			}
			count := helpers.MustAtoi(n)
			if count > counts[cube] {
				counts[cube] = count
			}
		}
	}

	return counts
}

func main() {
	lineRe := regexp.MustCompile(`^Game (\d+):\s+(.*)$`)

	total := 0
	powerTotal := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		matches := lineRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalln("Failed to match line", line)
		}
		gameNum := helpers.MustAtoi(matches[1])
		draws := strings.Split(matches[2], `; `)

		counts := minCubes(draws)
		if counts["red"] <= redMax && counts["green"] <= greenMax && counts["blue"] <= blueMax {
			fmt.Printf("Game %d possible\n", gameNum)
			total += gameNum
		}

		gamePower := counts["red"] * counts["green"] * counts["blue"]
		fmt.Printf("Game %d: power %d\n", gameNum, gamePower)
		powerTotal += gamePower
	})
	fmt.Println("Total:", total)
	fmt.Println("Power Total:", powerTotal)
}
