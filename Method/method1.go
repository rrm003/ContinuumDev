//go methods are the functions whose scope is narrowed to a specific type
//method recivers is an extra parameter used to specify host type

package main

import "fmt"

type gallon float64

func (g gallon) quart() gallon {
	return g * 4
}

func main() {
	//quart()	produce an error
	g := gallon(10) // variable of type method reciver
	//Access the quart() method using variable 'g'
	fmt.Println(g.quart())
}
