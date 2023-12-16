package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

func Hash(input string) int {
	value := 0
	for _, r := range input {
		value += int(r)
		value *= 17
		value = value % 256
	}
	return value
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	instructions := helpers.SplitCSV(string(input))
	hashTotal := 0
	for _, inst := range instructions {
		hash := Hash(inst)
		fmt.Printf("Instruction %s, hash:%d\n", inst, hash)
		hashTotal += hash
	}
	fmt.Println("Hash total:", hashTotal)
}
