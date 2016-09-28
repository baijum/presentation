package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	var v1 []int
	start1 := time.Now()
	for i := 0; i < 200000000; i++ {
		v1 = append(v1, i)
	}
	elapsed1 := time.Since(start1)
	fmt.Println(elapsed1, len(v1), cap(v1))

	var v2 []int
	start2 := time.Now()
	v2 = make([]int, 0, 200000000)
	for i := 0; i < 200000000; i++ {
		v2 = append(v2, i)
	}
	elapsed2 := time.Since(start2)
	fmt.Println(elapsed2, len(v2), cap(v2))
	// END OMIT
}
