package main

import "fmt"

func main() {
	// START OMIT
	i := 0
	for { // HL
		if i >= 5 {
			break
		}
		i++
		fmt.Println("Hello")
	}
	// END OMIT
}
