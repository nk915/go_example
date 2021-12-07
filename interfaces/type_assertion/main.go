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

// Skin ...
type Skin interface {
	Color() float64
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
	var s1 Shape = Cube{3}
	c := s1.(Cube)
	fmt.Println("volume of interface of type Shape is", c.Area())
	fmt.Println("volume of interface of type Object is", c.Volume())

	var s2 Shape = Cube{3}
	value1, ok1 := s2.(Object)
	fmt.Printf("dynamic value of Shape s with value %v implements interface object? %v\n", value1, ok1)
	value2, ok2 := s2.(Skin)
	fmt.Printf("dynamic value of Shape s with value %v implements interface object? %v\n", value2, ok2)
}

// --> Result
// volume of interface of type Shape is 54
// volume of interface of type Object is 27
// dynamic value of Shape s with value {3} implements interface object? true
// dynamic value of Shape s with value <nil> implements interface object? false
