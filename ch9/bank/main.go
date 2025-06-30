// Concurrency-safe bank with one account
package main

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is CONFINED to teller goroutine

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// Do nothing
		}
	}
}

func main() {
	done := make(chan struct{})
	go teller() // start the monitor goroutine
	Deposit(100)
	go func() { // consume with goroutine: use a channel to check
		Deposit(800)
		done <- struct{}{}
	}()
	Deposit(100)
	<-done
	close(done)
	fmt.Println(Balance()) // 1000
}
