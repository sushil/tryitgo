package main

import "fmt"

// will this deadlock
func main() {

	bc := make(chan bool)

	go func() {
		for {

		}
	}()

	fmt.Println("waiting")

	<-bc

	fmt.Println("done")

}
