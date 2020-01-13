package main

import "fmt"

func messages() chan int {

	ch := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			ch <- i // Send 'i' to channel 'ch'.
			// time.Sleep(1 * time.Millisecond)
		}

	}()

	return ch
}

// will this deadlock ?
// what if all the time.Sleep(..) calls are uncommented
func main() {

	ch := messages()

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	// time.Sleep(8 * time.Millisecond)
	fmt.Printf("done")

}
