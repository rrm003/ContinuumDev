package main

import "fmt"

func main() {
	for {
		var x int
		fmt.Scanf("%d", &x)
		if x%2 == 0{
			fmt.Println("Even")
		} else{
			fmt.Println("Odd")
		}
	}
}