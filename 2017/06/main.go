package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"os"

	"github.com/alext/aoc/helpers"
)

type MemoryBanks []int

func (m MemoryBanks) Checksum() uint32 {
	hash := crc32.NewIEEE()
	fmt.Fprintf(hash, "%v", m)
	return hash.Sum32()
}

func (m MemoryBanks) Redistribute() {
	max, maxIndex := 0, 0
	for i, b := range m {
		if b > max {
			max = b
			maxIndex = i
		}
	}
	m[maxIndex] = 0
	for i := maxIndex + 1; max > 0; {
		if i >= len(m) {
			i = 0
		}
		m[i]++
		i++
		max--
	}
}

func main() {
	var m MemoryBanks
	helpers.ScanWrapper(os.Stdin, bufio.ScanWords, func(word string) {
		m = append(m, helpers.MustAtoi(word))
	})
	seen := make(map[uint32]bool)
	seen[m.Checksum()] = true

	for count := 1; true; count++ {
		m.Redistribute()
		c := m.Checksum()
		if seen[c] {
			fmt.Printf("Duplicate configuration after %d cycles\n", count)
			break
		}
		seen[c] = true
	}

}
