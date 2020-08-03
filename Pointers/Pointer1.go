package main

import "fmt"

func main() {
	var arr = [5]int{10, 20, 30, 40, 50}
	var ptr *[5]int
	ptr = &arr
	fmt.Println(ptr[0])
}
