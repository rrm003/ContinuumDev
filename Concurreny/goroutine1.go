package main

import (
	"fmt"
)

func A(i int) {
	fmt.Println("I am : A", i)
}

func B(i int) {
	fmt.Println("I am : B", i)
}

func main() {
	for i := 0; i < 20; i++ {
		go A(i)
		go B(i)
	}
	fmt.Println("..closing main()")
}
