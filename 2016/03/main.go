package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Triangle []int

func (t Triangle) Possible() bool {
	sort.Ints(t)
	return t[0]+t[1] > t[2]
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle [%d, %d, %d]", t[0], t[1], t[2])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	possibleCount := 0

	var points [3][3]int
	index := 0

	for scanner.Scan() {
		_, err := fmt.Sscan(scanner.Text(), &points[index][0], &points[index][1], &points[index][2])
		if err != nil {
			log.Fatal(err)
		}
		index += 1
		if index == 3 {
			for i := 0; i < 3; i++ {
				t := Triangle{points[0][i], points[1][i], points[2][i]}
				fmt.Println(t)
				if t.Possible() {
					possibleCount += 1
				} else {
					fmt.Println("  impossible")
				}
			}
			index = 0
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Possible Count:", possibleCount)
}
