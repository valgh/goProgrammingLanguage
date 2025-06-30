// Example of an in-place slice algorithm
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // "one", "three"
	fmt.Printf("%q\n", data)           // "one", "three", "three" ---> that's why we usually re-assing the value to data!

}

// nonempty returns a slice holding only the non-empty strings
// it modifies the underlying array during the call
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
