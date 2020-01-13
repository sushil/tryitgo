package main

import (
	"encoding/json"
	"fmt"
)

type CountryAge struct {
	Country string `json:"country"`
	Age     int    `json:"age"`
}

type CopyDetails struct {
	Age      int    `json:"age,omitempty"`
	NickName string `json:"nick_name,omitempty"`
}

type NameCountry struct {
	Name    string     `json:"name"`
	Details CountryAge `json:"details"`
}

type Copy struct {
	Name    string      `json:"name"`
	Details CopyDetails `json:"details,omitempty"`
}

func main() {

	//source := `{
	//	"name": "source_name",
	//	"details": {
	//		"country": "USA",
	//		"age": 23,
	//		"nick_name": { "nn": 23 }
	//	}
	//}`

	source := `{
		"name": "source_name",
		"details": {
			"country_1": "USA",
			"age_": 23
		}
	}`

	var copy Copy
	e := json.Unmarshal([]byte(source), &copy)
	if e != nil {
		println(e.Error())
	}

	fmt.Printf("%v\n", copy)
	b, _ := json.MarshalIndent(copy, " ", " ")
	fmt.Printf("%s\n", b)

}
