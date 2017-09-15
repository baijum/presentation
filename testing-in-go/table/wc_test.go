package wc

import "testing"

func TestWordCount(t *testing.T) {
	testCases := []struct {
		s string
		c int
	}{
		{"hello", 1},
		{"hello world", 2},
		{"	twinkle twinkle little star	", 4},
	}
	for _, tc := range testCases {
		if r := WordCount(tc.s); r != tc.c {
			t.Errorf("Word count for '%s' should be %d. Result is %d", tc.s, tc.c, r)
		}
	}
}
