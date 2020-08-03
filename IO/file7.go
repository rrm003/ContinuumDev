package main

import (
	"fmt"
	"os"
)

type person struct {
	name   string
	age    int
	height float64
}

func main() {
	f, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	person1 := []person{{"rrm", 22, 12.01}, {"ajp", 19, 91.101}, {"rjt", 10, 11.1}, {"psg", 20, 10.01}}
	for _, p := range person1 {
		fmt.Fprintf(f, "%s \t %d \t %f \n", p.name, p.age, p.height)
	}
}
