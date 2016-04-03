package main

import "fmt"

// START OMIT
func sum(a int, b int) int {
	return a + b
}

func main() {
	v := sum(2, 3)
	fmt.Println(v)
}

// END OMIT
