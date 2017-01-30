package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/alext/aoc/helpers"
)

const MaxAddress = 4294967295

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

type Blacklist [][2]int64

func (b Blacklist) Len() int      { return len(b) }
func (b Blacklist) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b Blacklist) Less(i, j int) bool {
	return b[i][0] < b[j][0]
}

func (b Blacklist) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	for _, r := range b {
		fmt.Fprintf(&out, " %d-%d ", r[0], r[1])
	}
	out.WriteString("]")
	return out.String()
}

func (b *Blacklist) AddRange(start, end int64) {
	for i, r := range *b {
		if !(r[0] > end+1 || r[1] < start-1) {
			// Overlap
			(*b)[i][0] = min((*b)[i][0], start)
			(*b)[i][1] = max((*b)[i][1], end)
			for i < len(*b)-1 && (*b)[i][1]+1 >= (*b)[i+1][0] {
				// Overlap multiple ranges
				(*b)[i][1] = max((*b)[i][1], (*b)[i+1][1])
				// Remove merged item
				*b = (*b)[:i+1+copy((*b)[i+1:], (*b)[i+1+1:])]
			}
			sort.Sort(b)
			return
		}
	}

	*b = append(*b, [2]int64{start, end})
	sort.Sort(b)
}

var ErrNoneAvailable = errors.New("No addresses available")

func (b Blacklist) FirstAvailable() (int64, error) {
	if b[0][0] > 0 {
		return 0, nil
	}

	if b[0][1] == MaxAddress {
		return 0, ErrNoneAvailable
	}
	return b[0][1] + 1, nil
}

func (b Blacklist) NumAvailable() int64 {
	if len(b) == 0 {
		return MaxAddress
	}
	total := b[0][0]
	for i, r := range b {
		if i == len(b)-1 {
			total += MaxAddress - r[1]
		} else {
			//fmt.Printf("end: %d, start: %d, avail: %d\n", r[1], b[i+1][0]-1, (b[i+1][0]-1)-r[1])
			total += (b[i+1][0] - 1) - r[1]
		}
	}

	return total
}

func main() {
	b := &Blacklist{}
	helpers.ScanLines(os.Stdin, func(line string) {
		var start, end int64
		_, err := fmt.Sscanf(line, "%d-%d", &start, &end)
		if err != nil {
			log.Fatalln("Mismatched line in input:", line, "Error:", err)
		}
		b.AddRange(start, end)
	})

	first, err := b.FirstAvailable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("First available address:", first)
	fmt.Println("Total available addresses:", b.NumAvailable())
}
