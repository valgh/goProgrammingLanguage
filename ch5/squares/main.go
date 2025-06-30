// returns a function that returns the next square number each time it is called.
// This is an example of anonymous function
package main

import "fmt"

func squares() func() int {
	var x int
	return func() int { // anonymous function!
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println("When assigning a function to a variable, it maintains the state of its local variables.")
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	fmt.Println(f()) // 25

	fmt.Println("If NOT assigned to a variable, it won't maintain state.")
	fmt.Println(squares()()) // 1
	fmt.Println(squares()()) // 1
	fmt.Println(squares()()) // 1
}
