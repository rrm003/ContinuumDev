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

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	x := <-c
	fmt.Println(x)
	go sum(a[len(a)/2:], c)
	x = <-c
	fmt.Println(x)

}
