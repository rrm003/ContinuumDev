package main

import "fmt"

//Name is not exported
type Name *string

//method reciver cannot be of type pointer or interface
// func (name Name) getName() string {
// 	returnval := *name
// 	return returnval
// }

func main() {
	name := "rrm003"
	ptr := Name(&name)
	fmt.Println(*ptr)
	//ptr.getName()
}
