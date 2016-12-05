package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"
)

const doorID = "abbhdwsy"

func main() {
	var code bytes.Buffer

	foundChars := 0
	for i := 1; ; i++ {
		input := doorID + strconv.Itoa(i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))
		if hash[0:5] == "00000" {
			code.WriteByte(hash[5])
			foundChars += 1
			if foundChars >= 8 {
				break
			}
		}
	}
	fmt.Println("Code:", code.String())
}
