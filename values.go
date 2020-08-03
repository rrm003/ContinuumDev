package main

import (
	"fmt"
)

func swap(a, b int16) (int16, int16) {
	return b, a
}

func concat(a ...string) string {
	var name string
	for _, n := range a {
		name = name + " " + n
	}
	return name
}

func main() {

	//boolean
	var a bool = false
	fmt.Println("Boolean a -> ", a)

	//integer
	var x, y int16 = 10, 20
	fmt.Printf("Integer before swap\n x : %d, y :%d\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("Integer after swap\n x : %d, y :%d\n", x, y)

	//string
	var firstname, middlename, lastname string = "Ram", "Ravidas", "Mankar"
	name := concat(firstname, middlename)
	fmt.Printf("My name%s\n", name)
	fullname := concat(firstname, middlename, lastname)
	fmt.Printf("My Full Name%s\n", fullname)

}
