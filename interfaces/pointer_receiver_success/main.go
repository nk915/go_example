package main

import "fmt"

// Shape ...
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rect ...
type Rect struct {
	width  float64
	height float64
}

// Area ...
func (r *Rect) Area() float64 {
	return r.width * r.height
}

// Perimeter ...
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := Rect{5.0, 4.0}
	var s Shape = &r
	area := s.Area()
	perimeter := s.Perimeter()
	fmt.Println("Area of Rectangle is", area)
	fmt.Println("Perimeter of Rectangle is", perimeter)
}

// --> Result
// Area of Rectangle is 20
// Perimeter of Rectangle is 18
