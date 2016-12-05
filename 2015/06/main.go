package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/alext/aoc/helpers"
)

var lights [1000][1000]int

var instrRe = regexp.MustCompile(`(turn (?:on|off)|toggle)\s+(\d+),(\d+)\s+through\s+(\d+),(\d+)`)

func main() {
	helpers.ScanLines(os.Stdin, func(line string) {
		matches := instrRe.FindStringSubmatch(line)
		if matches == nil {
			fmt.Println("Failed to match line:", line)
			return
		}
		x1, _ := strconv.Atoi(matches[2])
		y1, _ := strconv.Atoi(matches[3])
		x2, _ := strconv.Atoi(matches[4])
		y2, _ := strconv.Atoi(matches[5])

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				switch matches[1] {
				case "turn on":
					lights[i][j] += 1
				case "turn off":
					if lights[i][j] > 0 {
						lights[i][j] -= 1
					}
				case "toggle":
					lights[i][j] += 2
				}
			}
		}
	})

	totalBrightness := 0
	for i := 0; i < len(lights); i++ {
		for j := 0; j < len(lights[0]); j++ {
			totalBrightness += lights[i][j]
		}
	}
	fmt.Println("Total Brightness:", totalBrightness)
}
