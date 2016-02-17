package main

import (
	"fmt"
)

// START OMIT
type Sounder interface {
	Sound() string
}

type Dog struct{}

func (d Dog) Sound() string {
	return "Woof"
}

func AnimalSound(s Sounder) string {
	return s.Sound()
}

func main() {
	d := Dog{}
	fmt.Println(AnimalSound(d))
}

// END OMIT
