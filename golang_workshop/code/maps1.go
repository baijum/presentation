package main

import "fmt"

func main() {
	// START OMIT
	var fruitWeights = map[string]int{
		"Apple":  45,
		"Mango":  24,
		"Orange": 34,
	}
	fmt.Printf("%#v\n", fruitWeights)
	// END OMIT
}
