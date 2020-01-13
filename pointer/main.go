package main

import "fmt"

func main() {
	var iptr *int64
	i := int64(200)
	iptr = &i
	fmt.Printf("%d %q %s\n", iptr, iptr, iptr)
	fmt.Printf("%d %q %s", *iptr, *iptr, *iptr)
}
