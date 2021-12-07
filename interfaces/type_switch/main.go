package main

import (
	"fmt"
	"strings"
)

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int ", i)
	default:
		fmt.Println("i stored something else ", i)
	}
}

func main() {
	explain("this is string")
	explain(12)
	explain(true)
}

// --> Result
// i stored string  THIS IS STRING
// i stored int  12
// i stored something else  true
