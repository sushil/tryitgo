package main

import (
	"fmt"
	"testing"
)
import "time"

func TestQ(t *testing.T) {

	fmt.Println("Will this deadlock ?")

	insertQ := make(chan int)
	updateQ := make(chan int)
	stop := make(chan bool)

	go func() {
		for toInsert := range insertQ {
			t := time.Now()
			fmt.Printf("[%s] i-loop: read from insertQ [%d]\n", t, toInsert)
			fmt.Printf("[%s] i-loop: waiting to put in updateQ [%d]\n", t, toInsert)
			// time.Sleep(time.Second * 1)
			go func() { updateQ <- toInsert }()
			fmt.Printf("[%s] i-loop: added to updateQ [%d]\n", t, toInsert)
		}
	}()

	go func() {
		for toUpdate := range updateQ {
			t := time.Now()
			fmt.Printf("[%s] u-loop: read from updateQ [%d]\n", t, toUpdate)
			fmt.Printf("[%s] u-loop: waiting to put in insertQ [%d]\n", t, toUpdate)
			// time.Sleep(time.Second * 1)
			go func() { insertQ <- toUpdate }()
			fmt.Printf("[%s] u-loop: added to insertQ [%d]\n", t, toUpdate)
		}
	}()

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Printf("main: waiting to put in insertQ [%d]\n", i)
			insertQ <- i
			fmt.Printf("main: added to insertQ [%d]\n", i)
		}
	}()

	go func() {

		for {

		}
	}()

	//insertQ <- 100
	//insertQ <- 200
	////insertQ <- 300

	<-stop

}
