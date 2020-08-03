package main

import (
	"fmt"
)

//Iterating array
func main() {
	var a [20]int
	a = [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i := 0; i < len(a); i++ {
		fmt.Printf("index: %d\tvalue: %d\n", i, a[i])
	}

	for i, val := range a {
		fmt.Printf("index: %d\tvalue: %d\n", i, val)
	}

}
