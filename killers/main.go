package main

import (
	"log"
	"time"
)

func main() {
	server()
}

func server() {
	var c chan struct{}
	var d chan struct{}

	// possible to late initialize signalling channel
	// but reader needs to check till its not nil
	go func() {
		time.After(2 * time.Second)
		c = make(chan struct{})
		close(c)
	}()

	go func() {
		time.After(2 * time.Second)
		d = make(chan struct{})
		close(d)
	}()

L:
	for {

		// reading from nil channel causes deadlock
		if c == nil || d == nil {
			log.Println("killers not ready")
			continue
		}

		select {
		case <-d:
			log.Println("d kills it")
			break L
		case <-c:
			log.Println("c kills it")
			break L
		}
	}

	log.Printf("out of here")
}
