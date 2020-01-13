package main

import (
	"fmt"
	"time"
)

func main() {
	if err := runMe(); err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("\n%s", "Done.")
}

func runMe() (e error) {

	go func(e *error) {

		defer func() {
			if r := recover(); r != nil {
				*e = fmt.Errorf("Recovered in runMe: %s", r)
			}
		}()

		time.Sleep(1 * time.Second)
		panic("you should not panic in goroutine")
	}(&e)

	// let the go routine's write to error have an effect before we return
	time.Sleep(4 * time.Second)

	return e
}
