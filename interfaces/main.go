package main

import (
	"fmt"
	"math"
)

type shape interface{
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {

	width float64
	height float64
}

func (c *circle) area() float64{
	return c.radius * c.radius * math.Pi
}

func (r *rectangle) area() float64{
	return r.height * r.width
}

func someFunction(someShape shape) {
	fmt.Println("Area of the shape:")
	fmt.Println(someShape.area())
}

func main(){

	circle1 := circle{radius: 4.5}
	rectangle1 := rectangle{width: 4, height: 5}

	someFunction(&circle1)
	someFunction(&rectangle1)



}