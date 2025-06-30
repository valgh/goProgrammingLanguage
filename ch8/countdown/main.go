package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	ticker := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
			// Do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		default:
			// Do nothing
		}
	}
	launch()
	ticker.Stop()
}

func launch() {
	fmt.Println("FFFFFSSSSSSSS Launched!")
}
