package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
	perimeter() float64
}

// -----

// Circle struct
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
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

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y1)
	w := distance(r.x1, r.y1, r.x1, r.y2)
	return 2 * (l + w)
}

// -----

// MultiShape struct
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, shape := range m.shapes {
		area += shape.area()
	}
	return area
}

// -----

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(totalArea(&c, &r))

	m := MultiShape{[]Shape{&c, &r}}
	fmt.Println(m.area())
	newM := MultiShape{[]Shape{&c, &r /*, &m /*- when Shape interface didn't require perimeter() method */}}
	fmt.Println(newM.area())
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.area()
	}
	return area
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
