// Example taken from Golang wiki: https://github.com/golang/go/wiki/PanicAndRecover
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// START HELPERS OMIT

// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
	Index int    // The index into the space-separated list of words.
	Word  string // The word that generated the parse error.
	Error error  // The raw error that precipitated this error, if any.
}

// String returns a human-readable error message.
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg: error parsing %q as int", e.Word)
}

// END HELPERS OMIT

// START PARSE OMIT

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

// END PARSE OMIT

// START FIELDS2NUMBERS OMIT

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}

// END FIELDS2NUMBERS OMIT

// START MAIN OMIT
func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n  ", ex)
		nums, err := Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}

// END MAIN OMIT
