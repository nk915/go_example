package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rect ...
type Rect struct {
	width  float64
	height float64
}

// Area returns area of rect
func (r Rect) Area() float64 {
	return r.width * r.height
}

// Perimeter returns perimeter of rect
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle ...
type Circle struct {
	radius float64
}

// Area returns area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter returns perimeter of Circle
func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func main() {
	var s Shape
	fmt.Println("value of s is", s)
	fmt.Printf("type of s is %T\n", s)

	fmt.Println("--------------------------------------")
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Println("Shape := Rect{5.0, 4.0}")
	fmt.Println("value of s is", s)
	fmt.Printf("type of s is %T\n", s)
	fmt.Println("Area of rectangle s is", s.Area())
	fmt.Println("s == r is", s == r)

	fmt.Println("--------------------------------------")
	s = Circle{10}
	fmt.Println("Shape := Circle{10}")
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Printf("value of s is %0.2f\n", s.Area())
}
