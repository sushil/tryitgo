package main

import "time"

// will this deadlock
func main() {

	bc := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
	}()

	<-bc

}
