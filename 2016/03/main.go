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

	for scanner.Scan() {
		t := Triangle{0, 0, 0}
		_, err := fmt.Sscan(scanner.Text(), &t[0], &t[1], &t[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
		if t.Possible() {
			possibleCount += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Possible Count:", possibleCount)
}
