package main

import "fmt"

func main() {

	//Array Declaration
	var students = [5]string{"rrm", "ajp", "rjt", "psg"}

	for i, s := range students {
		fmt.Println(i, " : ", s)
	}

	//length of array -> len()
	fmt.Println("Lenght of Array :", len(students))
}
