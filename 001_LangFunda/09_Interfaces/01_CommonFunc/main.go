package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Shape interface {
	// Function signature for Area
	area() float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func display(shape Shape) {
	fmt.Printf("%T\n", shape)
	fmt.Println(shape)
	fmt.Println(shape.area())
}

func main() {

	c1 := Circle{2}
	r1 := Rectangle{5, 2}

	display(c1)
	display(r1)
}
