package helpers

import (
	"bufio"
	"io"
	"log"
)

func ScanLines(in io.Reader, lineProcessor func(string)) {
	scanWrapper(in, bufio.ScanLines, lineProcessor)
}

func ScanRunes(in io.Reader, runeProcessor func(string)) {
	scanWrapper(in, bufio.ScanRunes, runeProcessor)
}

func scanWrapper(in io.Reader, split bufio.SplitFunc, processor func(string)) {
	scanner := bufio.NewScanner(in)
	scanner.Split(split)
	for scanner.Scan() {
		processor(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
