package main

import "fmt"

func main() {
	var a interface{}
	var b interface{}
	var c float64
	a = 10.00
	b = 10.0
	c = 10.0

	fmt.Println(a == b)
	fmt.Println(b == c)
}
