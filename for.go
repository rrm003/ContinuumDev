package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		//fmt.Print(i)
		for j := 0; j < 10; j++ {

			if i <= 10 && i <= j {
				fmt.Print("*")
			} else if i > j+9 {
				fmt.Print("*")
			}

		}
		fmt.Println()
	}
}
