package main

import "fmt"

func main() {
	// START OMIT
	if money := 20000; money > 15000 { // HL
		fmt.Println("I am going to buy a car.")
	} else {
		fmt.Println("I am going to buy a bike.")
	}
	// can't use the variable `money` here
	// END OMIT
}
