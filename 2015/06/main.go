package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/alext/aoc/helpers"
)

var lights [1000][1000]bool

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
					lights[i][j] = true
				case "turn off":
					lights[i][j] = false
				case "toggle":
					lights[i][j] = !lights[i][j]
				}
			}
		}
	})

	onCount := 0
	for i := 0; i < len(lights); i++ {
		for j := 0; j < len(lights[0]); j++ {
			if lights[i][j] {
				onCount++
			}
		}
	}
	fmt.Println("On count:", onCount)
}
