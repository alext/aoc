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
