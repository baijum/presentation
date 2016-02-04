package main

import "fmt"

func Sum(s chan int) {
	sum := 0
	for i := 1; i < 50; i++ {
		if i%3 == 0 {
			sum = sum + i
		} else {
			if i%5 == 0 {
				sum = sum + i
			}
		}
	}
	s <- sum
}

func main() {
	t := make(chan int)
	go Sum(t)
	fmt.Println("Sum:", <-t)
}
