package main

import (
	"fmt"
)

// START OMIT
var c = make(chan int)
var msg string

func setMessage() {
	msg = "Hello, World!"
	c <- 0 // HL
}

func main() {
	go setMessage()
	<-c // HL
	fmt.Println(msg)
}

// END OMIT
