package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Node struct {
	Size  int
	Used  int
	Avail int
	X     int
	Y     int
}

func main() {
	var nodes []*Node
	helpers.ScanLines(os.Stdin, func(line string) {
		if strings.Contains(line, "df -h") || strings.Contains(line, "Filesystem") {
			// Header lines
			return
		}
		n := &Node{}
		_, err := fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT", &n.X, &n.Y, &n.Size, &n.Used, &n.Avail)
		if err != nil {
			log.Fatal(err)
		}
		nodes = append(nodes, n)
	})

	viable := 0
	for a := 0; a < len(nodes); a++ {
		if nodes[a].Used == 0 {
			continue
		}
		for b := 0; b < len(nodes); b++ {
			if a == b {
				continue
			}
			if nodes[b].Avail >= nodes[a].Used {
				viable++
			}

		}
	}

	fmt.Println("Viable pairs:", viable)
}
