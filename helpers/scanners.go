package helpers

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func ScanLines(in io.Reader, lineProcessor func(string)) {
	ScanWrapper(in, bufio.ScanLines, lineProcessor)
}

func ScanRunes(in io.Reader, runeProcessor func(string)) {
	ScanWrapper(in, bufio.ScanRunes, runeProcessor)
}

func StreamRunes(in io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		ScanWrapper(in, bufio.ScanRunes, func(r string) { ch <- r })
		close(ch)
	}()
	return ch
}

func ScanWrapper(in io.Reader, split bufio.SplitFunc, processor func(string)) {
	scanner := bufio.NewScanner(in)
	scanner.Split(split)
	for scanner.Scan() {
		processor(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func SplitCSV(input string) []string {
	if input == "" {
		return []string{}
	}
	parts := strings.Split(input, ",")
	for i, _ := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
