package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func calculateRow(previous []byte) []byte {
	row := make([]byte, len(previous))

	for i := 0; i < len(row); i++ {
		if i == 0 {
			if previous[i+1] == '^' {
				row[i] = '^'
				continue
			}

		} else if i == len(row)-1 {
			if previous[i-1] == '^' {
				row[i] = '^'
				continue
			}
		} else {
			if previous[i-1] != previous[i+1] {
				row[i] = '^'
				continue
			}
		}
		row[i] = '.'
	}
	return row
}

func main() {
	roomHeight := flag.Int("height", 40, "The height of the room")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	grid := make([][]byte, *roomHeight)
	grid[0] = line[0 : len(line)-1] // strip trailing newline

	for i := 1; i < len(grid); i++ {
		grid[i] = calculateRow(grid[i-1])
	}
	count := 0
	for _, row := range grid {
		for _, tile := range row {
			if tile == '.' {
				count++
			}
		}
	}

	fmt.Println("Safe tiles:", count)
}
