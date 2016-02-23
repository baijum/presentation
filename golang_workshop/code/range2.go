package main

import "fmt"

func main() {
	// START OMIT
	var fruits = map[string]int{
		"Apple":  45,
		"Mango":  24,
		"Orange": 34,
	}
	for key, value := range fruits { // HL
		fmt.Println(key, value)
	}
	// END OMIT
}
