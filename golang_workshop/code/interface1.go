package main

import (
	"fmt"
)

// START OMIT
type Sounder interface {
	Sound() string
}

type Dog struct{}

type Cat struct{}

func (d Dog) Sound() string {
	return "Woof"
}

func (c Cat) Sound() string {
	return "Meow"
}

func AnimalSound(s Sounder) string {
	return s.Sound()
}

func main() {
	d, c := Dog{}, Cat{}
	fmt.Println(AnimalSound(d))
	fmt.Println(AnimalSound(c))
}

// END OMIT
