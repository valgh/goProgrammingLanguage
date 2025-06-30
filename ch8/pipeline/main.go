package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(naturals, squares)

	// Printer must be in main, as it's where the values from last channel are expected
	printer(squares)
}

func counter(naturals chan<- int) {
	for x := range 100 {
		naturals <- x
	}
	close(naturals)
}

func squarer(naturals <-chan int, squares chan<- int) {
	for {
		x, ok := <-naturals
		if !ok {
			break // naturals channel was closed!
		}
		squares <- x * x
	}
	close(squares)
}

func printer(squares <-chan int) {
	for x := range squares {
		fmt.Println(x)
	}
}
