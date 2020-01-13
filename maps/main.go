package main

import "fmt"

func main() {
	m1 := map[string]string{"one": "1"}
	m2 := map[string]string{}

	fmt.Printf("%v (%p) %v (%p)\n", m1, &m1, m2, &m2)

	m2 = m1
	fmt.Printf("%v (%p) %v (%p)\n", m1, &m1, m2, &m2)

	// will this work ?
	var m3 map[string]string
	m3["something"] = "something"

	fmt.Printf("%v", m3)
}
