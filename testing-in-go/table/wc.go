package wc

import "strings"

// WordCount count number words in the given string
func WordCount(s string) int {
	l := strings.Fields(s)
	return len(l)
}
