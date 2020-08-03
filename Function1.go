package main

import "fmt"

func half(integer int) bool {

	flag := false
	if integer%2 == 0 {
		flag = true
	}
	return flag
}

func main() {

	for {
		var i = 0
		fmt.Scanln(&i)
		fmt.Println(i, " ", half(i))
	}
}
