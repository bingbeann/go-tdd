package iteration

func Repeat(s string, n int) string {
	repeated := ""
	for range n {
		repeated += s
	}

	return repeated
}
