package main

import (
	"fmt"
	"strings"
)

func main() {
	arrayPrint()
}

func arrayPrint() {
	//	strArray := []string{"A", "B", "C"}
	strArray := []string{}

	fmt.Println(strings.Join(strArray, ","))

}
