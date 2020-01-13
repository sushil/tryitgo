package main

import (
	"errors"
	"fmt"
)

var (
	err0 = fmt.Errorf("error type 1")
	err1 = errors.New("error type 1")
	err2 = fmt.Errorf("error type 1")
	err3 = errors.New("error type 1")
	err4 = errors.New("error type 1")
)

func main() {

	fmt.Printf("err1 and err2 are equal: %v\n", err0 == err2)
	fmt.Printf("err1 and err2 are equal: %v\n", err1 == err2)
	fmt.Printf("err3 and err4 are equal: %v\n", err3 == err4)

	fmt.Printf("err0 and err4 are equal: %v\n", err2 == returnErrf0())
	fmt.Printf("err1 and err4 are equal: %v\n", err4 == returnErrors1())

	fmt.Printf("err0 and err4 are equal: %v\n", err0 == returnErrf0())
	fmt.Printf("err1 and err4 are equal: %v\n", err1 == returnErrors1())
}

func returnErrf0() error {
	return err0
}

func returnErrors1() error {
	return err1
}
