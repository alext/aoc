package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Room struct {
	encName  string
	sectorID int
	checksum string
}

func (r Room) String() string {
	return r.encName + " " + strconv.Itoa(r.sectorID) + " [" + r.checksum + "]"
}

func (r Room) Real() bool {
	return checksum(r.encName) == r.checksum
}

func (r Room) Name() string {
	return decrypt(r.encName, r.sectorID)
}

func checksum(input string) string {
	var counts helpers.LetterCounts
	for _, r := range input {
		if 'a' <= r && r <= 'z' {
			counts.Count(r)
		}
	}
	sort.Sort(&counts)
	chars := make([]byte, 0, 5)
	for _, i := range counts[0:5] {
		chars = append(chars, byte(i.Letter))
	}
	return string(chars)
}

func decrypt(input string, rot int) string {
	var output bytes.Buffer
	offset := rot % 26
	for _, r := range input {
		if 'a' <= r && r <= 'z' {
			r += rune(offset)
			if r > 'z' {
				r -= 26
			}
			output.WriteRune(r)
		} else if r == '-' {
			output.WriteRune(' ')
		}
	}
	return output.String()
}

var roomRegex = regexp.MustCompile(`([a-z-]+)-(\d+)\[([a-z]+)\]`)

func main() {
	realRooms := 0
	sectorIDSum := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		matches := roomRegex.FindStringSubmatch(line)
		if matches == nil || len(matches) < 4 {
			log.Println("Failed to match line:", line)
			return
		}
		sectorID, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal("Error parsing sectorID for", line, err)
		}
		r := Room{matches[1], sectorID, matches[3]}
		if r.Real() {
			realRooms += 1
			sectorIDSum += r.sectorID
			name := r.Name()
			if strings.Contains(name, "north") {
				fmt.Println(name, r.sectorID)
			}
		}
	})
	fmt.Println("Total real rooms:", realRooms)
	fmt.Println("Sum o their sectorIDs:", sectorIDSum)
}
