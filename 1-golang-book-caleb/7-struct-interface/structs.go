package main

import (
	"fmt"
	"math"
)

// Circle struct
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// -----

// Rectangle struct
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y1)
	w := distance(r.x1, r.y1, r.x1, r.y2)
	return l * w
}

// -----

func main() {

	fmt.Println("Declaration and Initialization:\n")

	//var c Circle
	//c := new(Circle)
	//c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)

	// -----
	fmt.Println()
	// -----

	fmt.Println("Functions that use structs:\n")

	fmt.Println(circleAreaUsual(c))
	fmt.Println(c) // there are no changes of c.x = 10

	fmt.Println(circleAreaPtr(&c))
	fmt.Println(c) // changes are affected

	// -----
	fmt.Println()
	// -----

	fmt.Println("Methods:\n")

	fmt.Println("Circle's area:", c.area())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Rectangle struct's area:", r.area())
}

func circleAreaUsual(c Circle) float64 {
	c.x = 10
	return math.Pi * c.r * c.r
}

func circleAreaPtr(c *Circle) float64 {
	c.x = 10
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
