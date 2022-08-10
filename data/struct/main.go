package main

import (
	"fmt"
	"reflect"
)

type A struct {
	a string
	b string
}

type B struct {
	a string
	b string
	c string
}

func main() {

	a := A{
		a: "A",
		b: "B",
	}

	b := A(a)
	c := a

	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))

}
