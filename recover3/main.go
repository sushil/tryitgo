package main

import (
	"fmt"
	"time"
)

func main() {
	errSignal := make(chan error)
	doneSignal := make(chan struct{})

	go func() {
		runMe(errSignal)
		// some work
		time.Sleep(3 * time.Second)
		doneSignal <- struct{}{}
	}()

	select {
	case e := <-errSignal:
		fmt.Printf("Error: %s", e)
	case <-doneSignal:
		fmt.Printf("\n%s", "Done.")
	}

}

func runMe(e chan error) {
	go func() {

		defer func() {
			if r := recover(); r != nil {
				e <- fmt.Errorf("Recovered in runMe: %s", r)
			}
		}()

		time.Sleep(1 * time.Second)
		panic("you should not panic in goroutine")
	}()
}
