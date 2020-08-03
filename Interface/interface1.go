package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (rec *rectangle) area() float64 {
	return rec.length * rec.breadth
}

func (rec *rectangle) perimeter() float64 {
	return (rec.length + rec.breadth) * 2
}

func (sq *square) area() float64 {
	return sq.side * sq.side
}

func (sq *square) perimeter() float64 {
	return sq.side * 4
}

func (cr *circle) area() float64 {
	return math.Pi * cr.radius * cr.radius
}

func (cr *circle) perimeter() float64 {
	return 2 * math.Pi * cr.radius
}

func main() {
	rec := rectangle(rectangle{12, 34})
	fmt.Println("Area : ", rec.area())
	fmt.Println("Perimeter : ", rec.perimeter())

	sq := square(square{66})
	fmt.Println("Area : ", sq.area())
	fmt.Println("Perimeter : ", sq.perimeter())

	cr := circle(circle{10})
	fmt.Println("Area : ", cr.area())
	fmt.Println("Perimeter : ", cr.perimeter())

}
