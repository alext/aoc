package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

func distanceMoved(time, holdTime int) int {
	return (time - holdTime) * holdTime
}

func numberOfWays(time, bestDistance int) int {
	lowestTime := 0
	for t := 0; t < time; t++ {
		if distanceMoved(time, t) > bestDistance {
			lowestTime = t
			break
		}
	}
	if lowestTime == 0 {
		// No options beat the best
		return 0
	}
	highestTime := 0
	for t := time; t >= 0; t-- {
		if distanceMoved(time, t) > bestDistance {
			highestTime = t
			break
		}
	}
	return highestTime - (lowestTime - 1)
}

func main() {
	input := bufio.NewReader(os.Stdin)
	timesLine, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	distancesLine, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var times []int
	var distances []int
	for _, n := range strings.Fields(timesLine)[1:] {
		times = append(times, helpers.MustAtoi(n))
	}
	for _, n := range strings.Fields(distancesLine)[1:] {
		distances = append(distances, helpers.MustAtoi(n))
	}

	fmt.Println(times)
	fmt.Println(distances)

	product := 1
	for i := range times {
		n := numberOfWays(times[i], distances[i])
		fmt.Printf("Time: %d, Distance: %d, ways: %d\n", times[i], distances[i], n)
		if n > 0 {
			product *= n
		}
	}
	fmt.Println("Product:", product)
}
