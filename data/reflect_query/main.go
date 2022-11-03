package main

import (
	"fmt"
	"reflect"
)

type MyType struct {
	A string `mytag:"ida"`
	B string `mytag:"idb"`
	C string `mytag:"-"`
}

func main() {
	//fmt.Printf("Hello World.\n")
	typeField()
}

func typeField() {
	m := MyType{"data-a", "data-b", "dada"}
	t := reflect.ValueOf(m)

	for i := 0; i < t.NumField(); i++ {
		v, ok := t.Type().Field(i).Tag.Lookup("mytag")
		if ok {
			fmt.Printf("tag(%s) field(%s) data(%s) \n", v, t.Type().Field(i).Name, t.Field(i))
		}
	}
}
