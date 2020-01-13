package main

import (
	"fmt"
	"time"
)

// will this deadlock
func main() {

	bc := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		bc <- true
	}()

	go func() {
		<-bc
	}()

	// what if this is uncommented
	// time.Sleep(2 * time.Second)
	<-bc

	fmt.Println("done")

}
