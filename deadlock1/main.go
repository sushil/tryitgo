package main

import (
	"fmt"
	"time"
)

func messages() chan int {

	fmt.Println("Will this deadlock ?")

	ch := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			ch <- i // Send 'i' to channel 'ch'.
			time.Sleep(2 * time.Second)
		}
		close(ch)
	}()

	return ch
}

func main() {

	ch := messages()
	closer(ch)

	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("done")
}

func closer(ch chan int) {
	go func() {
		fmt.Println("press enter to stop")
		var ip rune
		fmt.Scanf("%v", &ip)
		close(ch)
	}()
}
