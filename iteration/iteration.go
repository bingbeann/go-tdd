package iteration

import "strings"

func Repeat(s string, n int) string {
	var sb strings.Builder

	for range n {
		sb.WriteString(s)
	}

	return sb.String()
}
