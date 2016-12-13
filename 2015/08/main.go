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

func processEscape(in io.RuneReader) rune {
	r, _, err := in.ReadRune()
	if err != nil {
		log.Fatalln("Error processing escape:", err)
	}
	switch r {
	case '"', '\\':
		return r
	case 'x':
		digit1, _, err := in.ReadRune()
		if err != nil {
			log.Fatalln("Error reading:", err)
		}
		digit2, _, err := in.ReadRune()
		if err != nil {
			log.Fatalln("Error reading:", err)
		}
		out, err := hex.DecodeString(fmt.Sprintf("%c%c", digit1, digit2))
		if err != nil {
			log.Fatalln("Error decoding hex:", err)
		}
		return rune(out[0])
	default:
		log.Fatalf("Unrecognised escape sequence \\%c", r)
	}
	return ' ' // Never reached
}

func unescape(input string) string {
	in := bytes.NewBufferString(input)
	r, _, err := in.ReadRune()
	if err != nil || r != '"' {
		log.Fatalln("Line starts without quote, or read error:", err)
	}

	var out bytes.Buffer
	for r, _, err = in.ReadRune(); err == nil; r, _, err = in.ReadRune() {
		switch r {
		case '\\':
			_, err = out.WriteRune(processEscape(in))
			if err != nil {
				log.Fatalln("Error writing to buffer:", err)
			}
		case '"':
			break
		case ' ', '\n':
			// Ignore
		default:
			_, err = out.WriteRune(r)
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

func main() {
	inChars := 0
	outChars := 0
	helpers.ScanLines(os.Stdin, func(line string) {
		inChars += len(line)
		out := unescape(line)
		outChars += len(out)
		fmt.Println("In:", line, "Out:", out)
	})

	fmt.Println("Input characters:", inChars)
	fmt.Println("Output characters:", outChars)
	fmt.Println("Difference:", inChars-outChars)
}
