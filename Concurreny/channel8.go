package main

import "fmt"

func sum(x []int, c chan int) {
	var total int
	for _, val := range x {
		total = total + val
	}
	c <- total
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	c := make(chan int, 2)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}
}
