package main

import "fmt"

func zero(x *int) {
	*x = 0
	fmt.Println("In zero func")
}

func main() {
	x := 5
	fmt.Println("x : ", x)
	zero(&x)
	fmt.Println("x : ", x)
}
