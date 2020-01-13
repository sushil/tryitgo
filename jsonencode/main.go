package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	w := os.Stdout
	enc := json.NewEncoder(w)
	w.WriteString("[\n ")
	enc.Encode(Person{"bob", 12})
	w.WriteString(",")
	enc.Encode(Person{"alice", 14})
	w.WriteString("]")
}
