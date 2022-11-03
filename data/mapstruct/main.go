package main

import "fmt"

type test struct {
	a string
	b string
}

func main() {
	//mapPointer()
	//mapStruct()
	mapPointerList()
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
		mapTest[key] = &test{a: "a"}
	}

	if mapTest[key] == nil {
		fmt.Printf("nil \n")
	} else {
		fmt.Printf("not nil \n")
	}

	fmt.Printf("%v\n", mapTest[key])
}

func mapPointerList() {
	mapTest := make(map[string]*test)
	testList := []test{}
	key := "key"

	if mapTest[key] == nil {
		fmt.Printf("insert \n")
		mapTest[key] = &test{a: "a"}
		testList = append(testList, *mapTest[key])
	}

	if mapTest[key] == nil {
		fmt.Printf("nil \n")
	} else {
		mapTest[key].a = "b"
		fmt.Printf("not nil \n")
	}

	fmt.Printf("%v\n", mapTest[key])
	fmt.Printf("%v\n", testList)
}
