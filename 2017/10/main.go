package main

import (
	"fmt"
	"strings"

	"github.com/alext/aoc/helpers"
)

const (
	listSize  = 256
	blockSize = 16
	input     = "106,16,254,226,55,2,1,166,177,247,93,0,255,228,60,36"
	rounds    = 64
)

var extraLengths = []int{17, 31, 73, 47, 23}

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

type knotHash struct {
	list circularSlice
	pos  int
	skip int
}

func New(size int) *knotHash {
	h := &knotHash{
		list: make(circularSlice, size),
	}
	for i := 0; i < size; i++ {
		h.list[i] = i
	}
	return h
}

func (h *knotHash) runRound(lengths []int) {
	for _, length := range lengths {
		h.list.reverseSection(h.pos, length)
		h.pos += (length + h.skip) % len(h.list)
		h.skip++
	}
}

func (h *knotHash) denseHash() []int {
	blocks := len(h.list) / blockSize
	result := make([]int, blocks)
	for b := 0; b < blocks; b++ {
		for i := b * blockSize; i < (b+1)*blockSize; i++ {
			result[b] = result[b] ^ h.list[i]
		}
	}
	return result
}

func hexString(input []int) string {
	var out strings.Builder
	for _, i := range input {
		fmt.Fprintf(&out, "%02x", i)
	}
	return out.String()
}

func main() {
	lengths := make([]int, 0)
	for _, s := range helpers.SplitCSV(input) {
		lengths = append(lengths, helpers.MustAtoi(s))
	}

	hash := New(listSize)
	hash.runRound(lengths)
	fmt.Printf("First 2 numbers: %d, %d. Product: %d\n", hash.list[0], hash.list[1], hash.list[0]*hash.list[1])

	lengths2 := make([]int, 0)
	helpers.ScanRunes(strings.NewReader(input), func(r string) {
		lengths2 = append(lengths2, int(r[0]))
	})
	lengths2 = append(lengths2, extraLengths...)
	hash2 := New(listSize)
	for i := 0; i < rounds; i++ {
		hash2.runRound(lengths2)
	}
	fmt.Printf("First 2 numbers: %d, %d. Product: %d\n", hash2.list[0], hash2.list[1], hash2.list[0]*hash2.list[1])
	fmt.Println(hexString(hash2.denseHash()))
}
