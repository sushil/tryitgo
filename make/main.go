package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	items := make([]interface{}, 3)
	items[0] = "first string"
	items[1] = 3.14

	fmt.Printf("len(items) == %d\n", len(items))
	for i := range items {
		fmt.Printf("item[%d] == %#v\n", i, items[i])
	}

	fmt.Println("-- setting data from unmarshal into items --")

	data := `["first", "second"]`
	if err := json.Unmarshal([]byte(data), &items); err != nil {
		panic(err)
	}
	fmt.Printf("len(items) == %d\n", len(items))
	for i := range items {
		fmt.Printf("items[%d] == %#v\n", i, items[i])
	}

	fmt.Println("-- setting data from unmarshal into items2 --")
	var items2 []interface{}
	if err := json.Unmarshal([]byte(data), &items2); err != nil {
		panic(err)
	}
	fmt.Printf("len(items2) == %d\n", len(items2))
	for i := range items {
		fmt.Printf("items2[%d] == %#v\n", i, items2[i])
	}
}
