package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
)

func growData(in []byte) []byte {
	buf := bytes.NewBuffer(in[:])
	buf.Grow(len(in) + 1)
	buf.WriteByte('0')
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] == '0' {
			buf.WriteByte('1')
		} else {
			buf.WriteByte('0')
		}
	}
	return buf.Bytes()
}

func generateData(initial []byte, length int) []byte {
	data := initial
	for len(data) < length {
		data = growData(data)
	}
	return data
}

func checksumStage(in io.Reader) *bytes.Buffer {
	var out bytes.Buffer
	buf := make([]byte, 2)
	for {
		n, err := in.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatalln("Error reading:", err)
		}
		if n == 2 {
			if buf[0] == buf[1] {
				out.WriteByte('1')
			} else {
				out.WriteByte('0')
			}
		} else {
			break
		}
	}
	return &out
}

func calculateChecksum(data io.Reader, length int) []byte {
	data = &io.LimitedReader{R: data, N: int64(length)}
	checksum := checksumStage(data)
	for checksum.Len()%2 == 0 {
		checksum = checksumStage(checksum)
	}
	return checksum.Bytes()
}

func main() {
	initial := flag.String("initial", "", "Initial data sequence")
	diskSize := flag.Int("disk-size", 20, "Size of disk to be filled")

	flag.Parse()

	data := generateData([]byte(*initial), *diskSize)
	checksum := calculateChecksum(bytes.NewReader(data), *diskSize)

	fmt.Println("Checksum:", string(checksum))
}
