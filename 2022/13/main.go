package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alext/aoc/helpers"
)

type Value struct {
	List   *List
	Number int
}

func (v *Value) String() string {
	if v.List != nil {
		return fmt.Sprint(v.List)
	} else {
		return strconv.Itoa(v.Number)
	}
}

func (v *Value) Cmp(other *Value) int {
	// -1  l < other
	//  0  equal
	//  1  l > other
	if v.List == nil && other.List == nil {
		// Both numbers
		switch {
		case v.Number < other.Number:
			return -1
		case v.Number > other.Number:
			return 1
		default:
			return 0
		}
	}
	vList := v.List
	if vList == nil {
		vList = &List{v}
	}
	oList := other.List
	if oList == nil {
		oList = &List{other}
	}
	return vList.Cmp(*oList)
}

type List []*Value

func (l *List) String() string {
	return fmt.Sprint(*l)
}

func (l List) Cmp(other List) int {
	// -1  l < other
	//  0  equal
	//  1  l > other
	for i, v := range l {
		if i >= len(other) {
			// l is longer
			return 1
		}
		res := v.Cmp(other[i])
		if res != 0 {
			// found a difference
			return res
		}
	}
	// End of l. Check if other has more items, or is equal
	if len(l) == len(other) {
		return 0
	}
	// l has fewer items than other,
	return -1
}

func (l *List) addNumber(digits string) {
	if digits == "" {
		return
	}
	*l = append(*l, &Value{Number: helpers.MustAtoi(digits)})
}

func ParseList(in <-chan string) *List {
	var l List
	digits := ""
	for chr := range in {
		switch chr {
		case "]":
			l.addNumber(digits)
			return &l
		case "[":
			l = append(l, &Value{List: ParseList(in)})
		case ",":
			l.addNumber(digits)
			digits = ""
		default: // digit
			digits = digits + chr
		}
	}
	panic("Unexpected end of input")
}

func consume(in <-chan string, expected string) {
	c := <-in
	if c != expected {
		log.Fatalf("expected %s, got %s", expected, c)
	}
}

func main() {
	in := helpers.StreamRunes(os.Stdin)
	index := 1
	indexSum := 0
	for {
		consume(in, "[")
		l1 := ParseList(in)
		consume(in, "\n")

		consume(in, "[")
		l2 := ParseList(in)
		consume(in, "\n")

		res := l1.Cmp(*l2)
		if res <= 0 {
			indexSum += index
		}

		c, ok := <-in
		if !ok {
			// end of input
			break
		}
		if c != "\n" {
			log.Fatal("Expect blank line, got:", c)
		}
		index++
	}

	fmt.Println("IndexSum", indexSum)
}
