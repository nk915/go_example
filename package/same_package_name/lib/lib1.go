package lib

import "fmt"

var (
	StrLib1 = "lib1.go string"
)

func PrintLib1() {
	println("PrintLib1 print")
}

func PrintLibBy2() {
	fmt.Println(StrLib2)
	PrintLib2()
}
