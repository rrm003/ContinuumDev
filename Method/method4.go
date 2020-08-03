package main

import "fmt"

type gallon float64

//reciving value -> this will only update the copy recived by the method
func (g gallon) half() {
	fmt.Println("In half")
	g = gallon(g * 0.5)
	fmt.Println("Exit half")
}

//reciveing pointer -> this will update the original value
func (g *gallon) double() {
	fmt.Println("In double")
	*g = gallon(*g * 2)
	fmt.Println("Exit Half")
}

func main() {
	var gal gallon = 5
	gal.half()
	fmt.Println("Half :", gal)
	gal.double()
	fmt.Println("Double :", gal)
}
