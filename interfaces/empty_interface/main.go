package main

import (
	"fmt"
)

// MyString ...
type MyString string

// Rect ...
type Rect struct {
	width  float64
	height float64
}

func explain(i interface{}) {
	fmt.Printf("Value given to explain function is of Type %T with value %v\n", i, i)
}

func main() {
	ms := MyString("Hello World")
	r := Rect{5.0, 4.0}
	explain(ms)
	explain(r)
}

// --> Result
// Value given to explain function is of Type main.MyString with value Hello World
// Value given to explain function is of Type main.Rect with value {5 4}
