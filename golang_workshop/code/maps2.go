package main

import "fmt"

func main() {
	// START OMIT
	var fruitWeights map[string]int
	fruitWeights = make(map[string]int)
	fruitWeights["Apple"] = 45
	fruitWeights["Mango"] = 24
	fruitWeights["Orange"] = 34
	fmt.Printf("%#v\n", fruitWeights)
	// END OMIT
}
