package main

import (
	"fmt"
)

//sorting array
func sort(arr [10]int) {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

func main() {

	var arr [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}

	sort(arr)
}
