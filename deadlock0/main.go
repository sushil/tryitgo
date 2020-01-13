package main

import "fmt"

func messages() chan int {

	fmt.Println("Will this deadlock ?")

	ch := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			ch <- i // Send 'i' to channel 'ch'.
		}

		// close(ch)
	}()

	return ch
}

func main() {

	ch := messages()

	// what will this print ?
	fmt.Println(len(ch))

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

}
