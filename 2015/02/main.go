package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Present []int

func NewPresent(w, h, l int) Present {
	p := Present{w, h, l}
	sort.Ints(p)
	return p
}

func (p Present) Area() int {
	return 3*p[0]*p[1] + 2*p[1]*p[2] + 2*p[0]*p[2]
}

func (p Present) Ribon() int {
	return 2*p[0] + 2*p[1] + p[0]*p[1]*p[2]
}

func (p Present) String() string {
	return fmt.Sprintf("Triangle [%d, %d, %d]", p[0], p[1], p[2])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalArea := 0
	totalRibbon := 0

	for scanner.Scan() {
		var dims [3]int
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &dims[0], &dims[1], &dims[2])
		if err != nil {
			log.Fatal(err)
		}
		p := NewPresent(dims[0], dims[1], dims[2])
		totalArea += p.Area()
		totalRibbon += p.Ribon()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total area:", totalArea)
	fmt.Println("Total ribon:", totalRibbon)
}
