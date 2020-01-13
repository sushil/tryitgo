package main

import "fmt"

type operator struct {
	semaphore chan bool
}

func (o operator) do() {
	o.semaphore <- true

	//defer func() {
	//	<-o.semaphore
	//}()

	fmt.Print("- ")
}

func main() {
	opr := operator{semaphore: make(chan bool, 2)}

	for {
		go func() { opr.do() }()
	}
}
