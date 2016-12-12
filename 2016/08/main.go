package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/alext/aoc/helpers"
)

const (
	screenWidth  = 50
	screenHeight = 6
)

type Screen [screenWidth][screenHeight]bool

func (s *Screen) String() string {
	var res bytes.Buffer
	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			if s[x][y] {
				res.WriteString("#")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

func (s *Screen) Rect(x, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			s[i][j] = true
		}
	}
}

func (s *Screen) Rotate(axis string, index, amount int) {
	switch axis {
	case "row":
		s.rotateRow(index, amount)
	case "column":
		s.rotateCol(index, amount)
	default:
		panic("Unrecognised axis: " + axis)
	}
}

func (s *Screen) rotateRow(row, amount int) {
	var newValues [screenWidth]bool
	for x := 0; x < screenWidth; x++ {
		newValues[x] = s[(screenWidth+x-amount)%screenWidth][row]
	}
	for x := 0; x < screenWidth; x++ {
		s[x][row] = newValues[x]
	}
}

func (s *Screen) rotateCol(col, amount int) {
	var newValues [screenHeight]bool
	for y := 0; y < screenHeight; y++ {
		newValues[y] = s[col][(screenHeight+y-amount)%screenHeight]
	}
	for y := 0; y < screenHeight; y++ {
		s[col][y] = newValues[y]
	}
}

var (
	rectInstr   = regexp.MustCompile(`rect (\d+)x(\d+)`)
	rotateInstr = regexp.MustCompile(`rotate (row|column) (?:x|y)=(\d+) by (\d+)`)
)

func main() {
	scr := new(Screen)
	helpers.ScanLines(os.Stdin, func(line string) {
		if matches := rectInstr.FindStringSubmatch(line); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			scr.Rect(x, y)
		} else if matches := rotateInstr.FindStringSubmatch(line); matches != nil {
			index, _ := strconv.Atoi(matches[2])
			amount, _ := strconv.Atoi(matches[3])
			scr.Rotate(matches[1], index, amount)
		} else {
			log.Fatal("Unrecognised instruction:", line)
		}
	})

	fmt.Println(scr)
	litPixels := 0
	for x := 0; x < screenWidth; x++ {
		for y := 0; y < screenHeight; y++ {
			if scr[x][y] {
				litPixels++
			}
		}
	}
	fmt.Println("Lit pixels:", litPixels)
}
