package main

import "fmt"

func rotate(a []int, x int) {
	for i := 0; i < x; i++ {
		temp := a[0]
		for i := 1; i < len(a); i++ {
			a[i-1] = a[i]
		}
		a[len(a)-1] = temp
	}
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	rotate(a[:], 5)
	fmt.Println(a)
}
