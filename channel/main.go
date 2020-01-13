package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan error)
	go func() {
		e := <-c
		fmt.Printf("error is null? %v\n", e == nil)
	}()
	close(c)
	time.Sleep(2 * time.Second)
}
