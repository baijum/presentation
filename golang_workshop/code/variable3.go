package main

import "fmt"

func main() {
	// START OMIT
	var i = 42             // int
	var s, b = "foo", true // string, bool
	fmt.Printf("%#v, %#v, %#v\n", i, s, b)
	// END OMIT
}
