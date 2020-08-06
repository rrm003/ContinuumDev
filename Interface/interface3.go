package main

import (
	"fmt"
)

//Add ...
//go:generate mockgen -source=interface3.go -destination=mock_interface3.go -package=main Add
type Add interface {
	add(interface{}) interface{}
}

func add(in interface{}) interface{} {

	fmt.Println(in)
	switch c := in.(type) {
	case []int:
		var res int
		for _, x := range c {
			res = res + x
		}
		var out interface{} = res
		return out
	case []string:
		var res string
		for _, x := range c {
			res = res + x
		}
		var out interface{} = res
		return out
	case []float64:
		var res float64
		for _, x := range c {
			res = res + x
		}
		var out interface{} = res
		return out
	}
	return nil
}

func main() {
	var i interface{}
	i = []int{2, 3}
	fmt.Println(add(i))
	i = []float64{12.4, 29.8}
	fmt.Println(add(i))
	i = []string{"rrm", "003"}
	fmt.Println(add(i))
}
