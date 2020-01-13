package main

import (
	"fmt"
	"time"
)

func main() {
	if err := runMe(); err != nil {
		fmt.Printf("Error: %s", err)
	}

	time.Sleep(6 * time.Second)
	fmt.Printf("%s", "Done.")
}

func runMe() (e error) {

	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("Recovered in runMe: %s", r)
		}
	}()

	go func(e error) {

		defer func() {
			if r := recover(); r != nil {
				e = fmt.Errorf("Recovered in runMe: %s", r)
			}
		}()

		time.Sleep(2 * time.Second)
		panic("you should not panic in goroutine")
	}(e)

	time.Sleep(4 * time.Second)
	panic("you should not panic")

	return nil
}
