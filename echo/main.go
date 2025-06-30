/*
A simple command-line program that echoes
all the arguments passed by the user.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// easiest, need to import strings
	//fmt.Println(strings.Join(os.Args[1:], " "))
	for ix, arg := range os.Args[1:] {
		fmt.Println(ix, arg)
	}
}
