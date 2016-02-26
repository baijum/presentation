package main

import (
	"fmt"
	"sync"
)

// START OMIT
var msg string
var wg sync.WaitGroup

func setMessage() {
	msg = "Hello, World!"
	wg.Done() // HL
}

func main() {
	wg.Add(1) // HL
	go setMessage()
	wg.Wait() // HL
	fmt.Println(msg)
}

// END OMIT
