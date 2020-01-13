package main

import "fmt"

func main() {

	fmt.Println("Will this deadlock ?")

	ch := make(chan int)

	go func() {
		for i := range ch {
			fmt.Print(i)
		}
		fmt.Println("Will this print?")
	}()

	for {
		ch <- 1
	}

	fmt.Println("Done")

}


type chanelOwner struct{
	c chan int
}

func newChanelOwner(){
	co := chanelOwner{}
	co.init()
	return co
}

func co()