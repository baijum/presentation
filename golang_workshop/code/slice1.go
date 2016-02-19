package main

import "fmt"

func main() {
	// START OMIT
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
	// END OMIT
}
