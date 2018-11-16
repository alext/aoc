package main

import (
	"fmt"

	"github.com/alext/aoc/helpers"
)

const (
	listSize = 256
	input    = "106,16,254,226,55,2,1,166,177,247,93,0,255,228,60,36"
)

//const (
//listSize = 5
//input    = "3, 4, 1, 5"
//)

type circularSlice []int

func (c circularSlice) swap(i, j int) {
	i = i % len(c)
	j = j % len(c)
	c[i], c[j] = c[j], c[i]
}

func (c circularSlice) reverseSection(pos, length int) {
	for i, j := pos, pos+length-1; i < j; i, j = i+1, j-1 {
		c.swap(i, j)
	}
}

func main() {
	lengths := make([]int, 0)
	for _, s := range helpers.SplitCSV(input) {
		lengths = append(lengths, helpers.MustAtoi(s))
	}

	list := make(circularSlice, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = i
	}

	pos, skip := 0, 0
	for _, length := range lengths {
		list.reverseSection(pos, length)
		pos += (length + skip) % listSize
		skip++
	}

	fmt.Printf("First 2 numbers: %d, %d. Product: %d\n", list[0], list[1], list[0]*list[1])
}
