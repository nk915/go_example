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

	//	a := A{
	//		a: "A",
	//		b: "B",
	//	}
	//
	//	b := A(a)
	//	c := a
	//
	//	fmt.Println(b, reflect.TypeOf(b))
	//	fmt.Println(c, reflect.TypeOf(c))

	data := []string{"one", "two", "three"}
	example(data)
	moredata := []int{1, 2, 3}
	example(moredata)

	structdata := []B{B{a: "A"}, B{a: "B"}}
	example(structdata)

}

func example(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	}
}
