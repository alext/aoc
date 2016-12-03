package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const secret = "iwrupvqb"

func main() {
	for i := 1; ; i++ {
		input := secret + strconv.Itoa(i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))
		if hash[0:6] == "000000" {
			fmt.Println("Found:", i)
			break
		}
	}
}
