package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	fmt.Println("length : ", len(ch))
	fmt.Println("capacity : ", cap(ch))
	//fmt.Println(ch)
	ch <- 1
	//ch <- 2
}
