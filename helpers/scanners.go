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

func StreamLines(in io.Reader) <-chan string {
	return stream(in, bufio.ScanLines)
}

func ScanRunes(in io.Reader, runeProcessor func(string)) {
	ScanWrapper(in, bufio.ScanRunes, runeProcessor)
}

func StreamRunes(in io.Reader) <-chan string {
	return stream(in, bufio.ScanRunes)
}

func stream(in io.Reader, split bufio.SplitFunc) <-chan string {
	ch := make(chan string)
	go func() {
		ScanWrapper(in, split, func(r string) { ch <- r })
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

func ScanGrid(in io.Reader, separator string) [][]string {
	var result [][]string
	for line := range StreamLines(in) {
		if line == "" {
			break
		}
		result = append(result, strings.Split(line, separator))
	}
	return result
}
