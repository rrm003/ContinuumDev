package main

import "fmt"

//Person is not exported
type Person struct {
	name   string
	age    int
	weight float32
}

func (p Person) setPerson() {
	fmt.Println(p.name, p.age, p.weight)
}

func main() {
	p := Person{"rrm", 22, 73}
	p.setPerson()
}
