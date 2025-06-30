package main

import "fmt"

func main() {
	x := 10
	toAdd := 15

	fmt.Printf("Initial values - x: %d, toAdd: %d\n", x, toAdd)

	add(x, toAdd)
	fmt.Printf("Value of X after add by value x+toAdd: %d\n", x) // Still 10
	addWithPointer(&x, toAdd)
	fmt.Printf("Value of X after add by pointer *x+toAdd: %d\n", x)                                                     // Now x = 25
	fmt.Printf("Value of function that returns a new int performing return x + toAdd: %d\n", addAndReturnNew(x, toAdd)) // x(25)+15=40
}

func add(x int, toAdd int) {
	x = x + toAdd
}

func addWithPointer(x *int, toAdd int) {
	*x = *x + toAdd
}

func addAndReturnNew(x int, toAdd int) int {
	return x + toAdd
}
