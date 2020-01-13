package main

import "fmt"

func main() {
	cats := []Cat{
		Cat{Name: "one", Age: 2},
		Cat{Name: "two", Age: 3},
	}

	makeOlder(cats)
	// what will this print
	fmt.Printf("%v\n", cats)

}

type Cat struct {
	Name string
	Age  int
}

func makeOlder(cats []Cat) {
	for i := range cats {
		cats[i].Age = cats[i].Age + 10
	}

	cats = append(cats, Cat{Name: "three", Age: 5})
}
