package main

import "fmt"

// START OMIT
func div(a, b int) (q, rem int) {
	q, rem = a/b, a%b
	return
}

func main() {
	v, r := div(5, 2)
	fmt.Println(v, r)
}

// END OMIT
