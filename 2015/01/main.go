package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	floor := 0

	for scanner.Scan() {
		switch t := scanner.Text(); t {
		case "(":
			floor += 1
		case ")":
			floor -= 1
		default:
			fmt.Println("Unexpected character in input:", t)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Final floor:", floor)
}
