package main

import "fmt"

func sum(slice []int) int {

	ans := 0
	for _, b := range slice {

		ans = ans + b
	}
	return ans
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	slice := arr[1:4]
	fmt.Println(sum(slice))
}
