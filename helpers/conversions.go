package helpers

import "strconv"

func MustAtoi(in string) int {
	value, err := strconv.Atoi(in)
	if err != nil {
		panic("Failed to parse number: " + in)
	}
	return value
}
