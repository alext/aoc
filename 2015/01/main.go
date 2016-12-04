package main

import (
	"fmt"
	"os"

	"github.com/alext/aoc/2015/helpers"
)

func main() {
	floor := 0
	charPos := 1

	helpers.ScanRunes(os.Stdin, func(t string) {
		switch t {
		case "(":
			floor += 1
		case ")":
			floor -= 1
		default:
			fmt.Println("Unexpected character in input:", t)
		}
		if floor < 0 {
			fmt.Println("Entered basement at position:", charPos)
			os.Exit(0)
		}
		charPos += 1
	})
	fmt.Println("Final floor:", floor)
}
