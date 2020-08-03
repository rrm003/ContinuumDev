package main

import "fmt"

func main() {
	ptr := new([2]int)
	*ptr = [2]int{10, 20}
	fmt.Println(*ptr)
}
