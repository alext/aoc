package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type readerByteReader interface {
	io.Reader
	io.ByteReader
}

func expandSequence(in readerByteReader) (int, error) {
	var length, count int
	_, err := fmt.Fscanf(in, "%dx%d)", &length, &count)
	if err != nil {
		log.Fatal("Error in Fscanf", err)
		return 0, err
	}

	var segment bytes.Buffer
	_, err = io.CopyN(&segment, in, int64(length))
	if err != nil {
		return 0, err
	}

	return segment.Len() * count, nil
}

func expandStream(in readerByteReader) (int, error) {
	outCount := 0
	for r, err := in.ReadByte(); err == nil; r, err = in.ReadByte() {
		switch r {
		case '(':
			seqSize, err := expandSequence(in)
			if err != nil {
				log.Fatal(err)
			}
			outCount += seqSize
		case ' ', '\n':
		default:
			outCount++
		}
	}

	return outCount, nil
}

func main() {
	outCount, err := expandStream(bufio.NewReader(os.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Output count:", outCount)
}
