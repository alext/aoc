package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Present []int

func (p Present) Area() int {
	sort.Ints(p)
	return 3*p[0]*p[1] + 2*p[1]*p[2] + 2*p[0]*p[2]
}

func (p Present) String() string {
	return fmt.Sprintf("Triangle [%d, %d, %d]", p[0], p[1], p[2])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalArea := 0

	for scanner.Scan() {
		var dims [3]int
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &dims[0], &dims[1], &dims[2])
		if err != nil {
			log.Fatal(err)
		}
		p := Present{dims[0], dims[1], dims[2]}
		totalArea += p.Area()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total area:", totalArea)
}
