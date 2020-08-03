package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("A : ", i)
		}
	}()

	for x := 0; x < 10; x++ {
		fmt.Println("Main ", <-ch)
	}
}
