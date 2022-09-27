package main

import "fmt"

type test struct {
	a string
	b string
}

func main() {
	//mapPointer()
	mapStruct()
}

func mapStruct() {
	mapTest := make(map[string]test)
	key := "key"

	if len(mapTest[key].a) == 0 {
		fmt.Printf("insert \n")
		mapTest[key] = test{a: "a"}
	}

	if len(mapTest[key].a) == 0 {
		fmt.Printf("nil \n")
	} else {
		fmt.Printf("not nil \n")
	}

	fmt.Printf("%v\n", mapTest[key])
}

func mapPointer() {
	mapTest := make(map[string]*test)
	key := "key"

	if mapTest[key] == nil {
		fmt.Printf("insert \n")
		mapTest[key] = &test{}
	}

	if mapTest[key] == nil {
		fmt.Printf("nil \n")
	} else {
		fmt.Printf("not nil \n")
	}

	fmt.Printf("%v\n", mapTest[key])
}
