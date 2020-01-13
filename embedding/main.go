package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID string `json:"id"`
	Details
}

type Details struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

const data = `{
	"id": "123",
	"first_name": "jane",
	"last_name": "doe"
}`

func main() {
	var p Person
	err := json.Unmarshal([]byte(data), &p)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", p)

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", b)
}
