package main

import (
	"fmt"

	"github.com/sdslabs/kiwi/stdkiwi"
)

func main() {
	store := stdkiwi.NewStore()

	str(store)
	hash(store)

}

func hash(store *stdkiwi.Store) {
	if err := store.AddKey("consul_hash", "hash"); err != nil {
		panic(err)
	}

	students := store.Hash("consul_hash") // assumes "students" key is of hash type

	if err := students.Insert("/namespace/aa", "111"); err != nil {
		panic(err)
	}
	if err := students.Insert("/namespace/bb", "222"); err != nil {
		panic(err)
	}
	if err := students.Insert("/namespace/bb", "bbb"); err != nil {
		panic(err)
	}

	if find, err := students.Get("/namespace/bb"); err != nil {
		fmt.Println("not found bb")
	} else {
		fmt.Println("find: ", find)
	}

	if find, err := students.Get("/namespace/cc"); err != nil {
		fmt.Println("not found cc")
	} else {
		fmt.Println("find: ", find)
	}

}

func str(store *stdkiwi.Store) {

	if err := store.AddKey("my_string", "str"); err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	myString := store.Str("my_string")

	if err := myString.Update("Hello, World!"); err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	str, err := myString.Get()
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	fmt.Println(str) // Hello, World!
}
