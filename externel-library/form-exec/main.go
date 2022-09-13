package main

import (
	"fmt"
	"log"

	//"github.com/go-playground/form/v4"
	"github.com/nk915/form"
)

// Address contains address information
type Address struct {
	Name  string
	Phone string
	Email []string
}

// User contains user information
type User struct {
	Name        string
	Age         uint8
	Site        string `form:",omitempty"`
	Gender      string
	Address     []Address
	Active      bool `form:"active"`
	MapExample  map[string]string
	NestedMap   map[string]map[string]string
	NestedArray [][]string
}

// use a single instance of Encoder, it caches struct info
var encoder *form.Encoder

func main() {
	encoder = form.NewEncoder()

	user := User{
		Name:   "joeybloggs",
		Age:    3,
		Gender: "Male",
		Address: []Address{
			{Name: "26 Here Blvd.", Phone: "9(999)999-9999", Email: []string{"nk915@aaa.com", "nk915@bbb.com"}},
			{Name: "26 There Blvd.", Phone: "1(111)111-1111"},
		},
		Active:      true,
		MapExample:  map[string]string{"key": "value"},
		NestedMap:   map[string]map[string]string{"key": {"key": "value"}},
		NestedArray: [][]string{{"value"}},
	}

	// must pass a pointer
	values, err := encoder.Encode(&user)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n\n", values)
	fmt.Printf("%s", values.Encode())
}
