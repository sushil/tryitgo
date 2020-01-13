package main

import (
	"fmt"
	"time"
)

// will this deadlock
func main() {

	bc := make(chan bool)

	go func() {
		bc <- true
	}()

	go func() {
		<-bc
		<-bc
	}()

	time.Sleep(2 * time.Second)

	<-bc

	fmt.Println("done")

}
