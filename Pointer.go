package main

import (
	"fmt"
)

func main() {

	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var p1, p2 *int //= &arr[0], &arr[9]

	//reverse
	for i := 0; i < len(arr)/2; i++ {
		p1 = &arr[i]
		p2 = &arr[len(arr)-i-1]
		fmt.Println(*p1, *p2)
	}

	fmt.Println(arr)
}
