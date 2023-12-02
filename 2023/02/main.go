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

func drawPossible(draw string) bool {
	cubes := strings.Split(draw, `, `)
	for _, cubeDraw := range cubes {
		n, cube, ok := strings.Cut(cubeDraw, ` `)
		if !ok {
			log.Fatalf("Invalid cubes %s in draw %s", cubeDraw, draw)
		}
		count := helpers.MustAtoi(n)
		switch cube {
		case "red":
			if count > redMax {
				return false
			}
		case "green":
			if count > greenMax {
				return false
			}
		case "blue":
			if count > blueMax {
				return false
			}
		default:
			log.Fatalf("Unexpected colour %s in draw %s", cube, draw)
		}
	}
	return true
}

func main() {
	lineRe := regexp.MustCompile(`^Game (\d+):\s+(.*)$`)

	total := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		matches := lineRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalln("Failed to match line", line)
		}
		gameNum := helpers.MustAtoi(matches[1])
		draws := strings.Split(matches[2], `; `)

		for _, draw := range draws {
			if !drawPossible(draw) {
				fmt.Printf("Game %d: draw %s not possible\n", gameNum, draw)
				return
			}
		}
		total += gameNum
	})
	fmt.Println("Total:", total)
}
