package main

import "fmt"

type Wrap struct {
	Color string
	Candy Candy
}

type Candy struct {
	Flavor string
}

func main() {
	wraps, candies := makeThem()

	fmt.Printf("Wrapped %v -- Candies %v\n", wraps, candies)
	fmt.Printf(`&wraps[0].Candy: %p -- vs -- &candies[0]: %p`, &wraps[0].Candy, &candies[0])

	candies[0].Flavor = `**yucky**`

	fmt.Printf("\n\nWrapped %v -- Candies %v\n", wraps, candies)
	fmt.Printf(`&wraps[0].Candy: %p -- vs -- &candies[0]: %p`, &wraps[0].Candy, &candies[0])
}

func makeThem() (wraps []Wrap, candies []Candy) {
	colors := []string{"red", "blue", "yellow"}
	flavors := []string{"sweet", "sour", "spicy"}

	for i, f := range flavors {
		candy := Candy{Flavor: f}
		wrap := Wrap{Color: colors[i], Candy: candy}

		candies = append(candies, candy)
		wraps = append(wraps, wrap)
	}

	return wraps, candies
}
