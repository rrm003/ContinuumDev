//Implementing Multiple Interface

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type polygon interface {
	perimeter() float64
}

type curve interface {
	circumference() float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

func (rec *rectangle) area() float64 {
	return rec.length * rec.breadth
}

func (cr *circle) area() float64 {
	return math.Pi * cr.radius * cr.radius
}

func (rec *rectangle) perimeter() float64 {
	return 2 * (rec.length + rec.breadth)
}

func (cr *circle) circumference() float64 {
	return 2 * math.Pi * cr.radius
}

func main() {
	rec := rectangle(rectangle{10, 20})
	fmt.Println("Rectangle : ")
	fmt.Println("Area : ", rec.area())
	fmt.Println("Permiter : ", rec.perimeter())

	cr := circle(circle{100})
	fmt.Println("Circle : ")
	fmt.Println("Area : ", cr.area())
	fmt.Println("Permiter : ", cr.circumference())
}
