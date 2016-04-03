package main

import "fmt"

// START OMIT
func div(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	v, r := div(5, 2)
	fmt.Println(v, r)
}

// END OMIT
