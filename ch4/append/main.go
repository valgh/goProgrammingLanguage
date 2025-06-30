package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d     cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	x = append(x, x...) //  built-in append function: can even append entire slices
	fmt.Printf("%d     cap=%d\t%v\n", 10, cap(x), x)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		// there is room to grow: extend the slice
		z = x[:zlen]
	} else {
		// insufficient space: allocate a new array, grow by doubling
		zcap := max(zlen, 2*len(x)) // built-in max function: assign to the variable the maximum between the two values
		z = make([]int, zlen, zcap)
		copy(z, x) // built-in copy function, copies x into z
	}

	z[len(x)] = y
	return z
}
