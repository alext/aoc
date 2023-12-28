package helpers

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LeastCommonMultiple(numbers ...int) int {
	if len(numbers) < 2 {
		return numbers[0]
	}
	a, b, rest := numbers[0], numbers[1], numbers[2:]
	res := a * b / GreatestCommonDivisor(a, b)
	for _, next := range rest {
		res = LeastCommonMultiple(res, next)
	}
	return res
}
