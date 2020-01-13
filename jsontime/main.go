package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Test struct {
	A time.Time `json:"a"`
	B JSONTime  `json:"b"`
}

// JSONTime is only a wrapper struct to have a custom marshaling behavior.
type JSONTime struct {
	time.Time
}

// MarshalJSON is the interface implementation of the json.Marshaler
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return json.Marshal(nil)
	}
	return t.Truncate(time.Second).MarshalJSON()
}

func main() {

	t := Test{
		A: time.Now(),
		B: JSONTime{time.Now()},
	}

	bytes, _ := json.MarshalIndent(&t, "", "    ")

	fmt.Println(string(bytes))
}