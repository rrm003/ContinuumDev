package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func B(ch chan int) {
	a := <-ch
	fmt.Println("B : ", a)
	a++
	ch <- a
	wg.Done()
}

func A(ch chan int) {
	a := <-ch
	fmt.Println("A : ", a)
	a++
	ch <- a
	wg.Done()
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go A(ch)
		go B(ch)
		wg.Wait()
	}
}
