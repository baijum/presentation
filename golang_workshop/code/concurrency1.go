package main

import (
	"fmt"
	"time"
)

// START OMIT
var msg string

func setMessage() {
	msg = "Hello, World!"
}

func main() {
	go setMessage() // HL
	time.Sleep(1 * time.Millisecond)
	fmt.Println(msg)
}

// END OMIT
