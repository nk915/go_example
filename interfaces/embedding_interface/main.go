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

// Material ...
type Material interface {
	Shape
	Object
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
	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Printf("dynamic type and value of m of static type Material is %T and %v\n", m, m)
	fmt.Printf("dynamic type and value of s of static type Shape is %T and %v\n", s, s)
	fmt.Printf("dynamic type and value of o of static type Object is %T and %v\n", o, o)
}

// --> Result
// dynamic type and value of m of static type Material is main.Cube and {3}
// dynamic type and value of s of static type Shape is main.Cube and {3}
// dynamic type and value of o of static type Object is main.Cube and {3}
