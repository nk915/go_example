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
type MyType2 struct {
	D string `mytag:"idd"`
}

type MyType3 struct {
	E     string `mytag:"ide"`
	type2 MyType2
}

func main() {
	//fmt.Printf("Hello World.\n")
	m := MyType{"data-a", "data-b", "dada"}
	m2 := MyType2{"data-d"}
	m3 := MyType3{E: "data-e", type2: m2}

	var interface_list []interface{}
	interface_list = append(interface_list, m, m2)

	for _, data := range interface_list {
		makeInventoryVars(data)
	}

	//typeField(m)
	fmt.Printf("--- \n")
	makeInventoryVars(m3)

}

func typeField(m MyType) {
	t := reflect.ValueOf(m)
	for i := 0; i < t.NumField(); i++ {
		v, ok := t.Type().Field(i).Tag.Lookup("mytag")
		if ok {
			fmt.Printf("tag(%s) field(%s) data(%s) \n", v, t.Type().Field(i).Name, t.Field(i))
		}
	}
}

func typeInterfaceField(data interface{}) {
	t := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		v, ok := t.Type().Field(i).Tag.Lookup("mytag")
		if ok {
			fmt.Printf("tag(%s) field(%s) data(%s) \n", v, t.Type().Field(i).Name, t.Field(i))
		}
	}
}

func makeInventoryVars(vars interface{}) error {

	ref := reflect.ValueOf(vars)
	inputTagVars(ref)
	//for i := 0; i < ref.NumField(); i++ {
	//	inputTagVars(ref.Field(i))
	//}

	return nil
}

func inputTagVars(ref reflect.Value) {
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		if field.Kind() == reflect.Struct {
			inputTagVars(field)
			continue
		}
		tag, ok := ref.Type().Field(i).Tag.Lookup("mytag")
		if ok {
			fmt.Printf("%s=%s\n", tag, ref.Field(i))
		}
	}
}
