package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const doorID = "abbhdwsy"

func main() {
	code := make([]byte, 8)
	for i := 0; i < len(code); i++ {
		code[i] = '_'
	}
	fmt.Print("Code: ", string(code))

	for i := 1; ; i++ {
		input := doorID + strconv.Itoa(i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))
		if hash[0:5] == "00000" {
			if '0' <= hash[5] && hash[5] <= '7' {
				pos := hash[5] - '0'
				if code[pos] == '_' {
					code[pos] = hash[6]
					fmt.Print("\rCode: ", string(code))
				}
				if !strings.Contains(string(code), "_") {
					break
				}
			}
		}
	}
	fmt.Println("")
}
