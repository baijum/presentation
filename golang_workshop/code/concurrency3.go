package main

import (
	"fmt"
	"sync"
)

// START OMIT
var l sync.Mutex
var msg string

func setMessage() {
	msg = "Hello, World!"
	l.Unlock()
}

func main() {
	l.Lock()
	go setMessage()
	l.Lock()
	fmt.Println(msg)
}

// END OMIT
