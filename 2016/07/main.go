package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type IPv7 struct {
	addressSegments   []string
	hypernetSequences []string
}

func Parse(input io.Reader) *IPv7 {
	var addr = &IPv7{}
	var segment bytes.Buffer
	helpers.ScanRunes(input, func(char string) {
		switch char {
		case "[":
			addr.addressSegments = append(addr.addressSegments, segment.String())
			segment.Reset()
		case "]":
			addr.hypernetSequences = append(addr.hypernetSequences, segment.String())
			segment.Reset()
		default:
			segment.WriteString(char)
		}
	})
	addr.addressSegments = append(addr.addressSegments, segment.String())
	return addr
}

func (i *IPv7) SupportsTLS() bool {
	for _, s := range i.hypernetSequences {
		if containsABBA(s) {
			return false
		}
	}
	for _, s := range i.addressSegments {
		if containsABBA(s) {
			return true
		}
	}
	return false
}

func (i *IPv7) SupportsSSL() bool {
	babs := make([]string, 0)
	for _, s := range i.addressSegments {
		for _, bab := range getBABs(s) {
			babs = append(babs, bab)
		}
	}
	if len(babs) == 0 {
		return false
	}
	for _, s := range i.hypernetSequences {
		for _, bab := range babs {
			if strings.Contains(s, bab) {
				return true
			}
		}
	}
	return false
}

func containsABBA(in string) bool {
	window := make([]rune, 0)
	for _, r := range in {
		if len(window) >= 4 {
			window = window[1:4]
		}
		window = append(window, r)
		if len(window) < 4 {
			continue
		}
		if window[0] == window[3] && window[1] == window[2] && window[0] != window[1] {
			return true
		}
	}
	return false
}

func getBABs(in string) []string {
	babs := make([]string, 0)
	window := make([]rune, 0)
	for _, r := range in {
		if len(window) >= 3 {
			window = window[1:3]
		}
		window = append(window, r)
		if len(window) < 3 {
			continue
		}
		if window[0] == window[2] && window[0] != window[1] {
			babs = append(babs, fmt.Sprintf("%c%c%c", window[1], window[0], window[1]))
		}
	}
	return babs
}

func main() {
	tlsAddresses := 0
	sslAddresses := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		addr := Parse(strings.NewReader(line))
		if addr.SupportsTLS() {
			tlsAddresses += 1
		}
		if addr.SupportsSSL() {
			sslAddresses += 1
		}
	})
	fmt.Println("TLS supporting addresses:", tlsAddresses)
	fmt.Println("SSL supporting addresses:", sslAddresses)
}
