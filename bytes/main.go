package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	c := cat{Name: "meo", Age: 5}
	b, _ := json.Marshal(c)
	s := string(b)
	fmt.Println(s)
}

type cat struct {
	Name string
	Age  int
}
