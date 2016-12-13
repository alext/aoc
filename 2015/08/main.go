package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

func processEscape(in *bytes.Buffer) byte {
	b, err := in.ReadByte()
	if err != nil {
		log.Fatalln("Error processing escape:", err)
	}
	switch b {
	case '"', '\\':
		return b
	case 'x':
		digits := in.Next(2)
		if len(digits) != 2 {
			log.Fatalln("Failed to read 2 digits for \\x escape sequence")
		}
		out, err := hex.DecodeString(string(digits))
		if err != nil {
			log.Fatalln("Error decoding hex:", err)
		}
		return out[0]
	default:
		log.Fatalf("Unrecognised escape sequence \\%c", b)
	}
	return ' ' // Never reached
}

func unescape(input string) string {
	in := bytes.NewBufferString(input)
	b, err := in.ReadByte()
	if err != nil || b != '"' {
		log.Fatalln("Line starts without quote, or read error:", err)
	}

	var out bytes.Buffer
	for b, err = in.ReadByte(); err == nil; b, err = in.ReadByte() {
		switch b {
		case '\\':
			err = out.WriteByte(processEscape(in))
			if err != nil {
				log.Fatalln("Error writing to buffer:", err)
			}
		case '"':
			break
		case ' ', '\n':
			// Ignore
		default:
			err = out.WriteByte(b)
			if err != nil {
				log.Fatalln("Error writing to buffer:", err)
			}
		}
	}
	if err != io.EOF {
		log.Fatalln("Error reading:", err)
	}

	return out.String()
}

func escape(input string) string {
	in := bytes.NewBufferString(input)
	out := bytes.NewBufferString(`"`)

	for b, err := in.ReadByte(); err == nil; b, err = in.ReadByte() {
		switch b {
		case '"', '\\':
			out.WriteByte('\\')
			fallthrough
		default:
			out.WriteByte(b)
		}
	}
	out.WriteByte('"')
	return out.String()
}

func main() {
	inChars := 0
	decodedChars := 0
	encodedChars := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		inChars += len(line)
		decodedChars += len(unescape(line))
		encodedChars += len(escape(line))
	})

	fmt.Println("Input characters:", inChars)
	fmt.Println("Decoded characters:", decodedChars)
	fmt.Println("Decoded difference:", inChars-decodedChars)
	fmt.Println("Encoded characters:", encodedChars)
	fmt.Println("Encoded difference:", encodedChars-inChars)
}
