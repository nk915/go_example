package main

import (
	"fmt"
)

// Shape ...
type Shape interface {
	Area() float64
}

// Object ...
type Object interface {
	Volume() float64
}

// Cube ...
type Cube struct {
	side float64
}

// Area ...
func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

// Volume ...
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{3}
	var s Shape = c
	var o Object = c
	fmt.Println("volume of interface of type Shape is", s.Area())
	fmt.Println("volume of interface of type Object is", o.Volume())
}

// --> Result
// volume of interface of type Shape is 54
// volume of interface of type Object is 27
