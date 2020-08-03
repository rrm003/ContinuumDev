package main

import "fmt"

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < len(a)/2; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	reverse(a[:])
	fmt.Println(a)
}
