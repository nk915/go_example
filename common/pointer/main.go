package main

import "fmt"

const (
	MAX_RANDOM_LEN = 32
)

func main() {

	//var test [MAX_RANDOM_LEN]byte
	test := make([]byte, MAX_RANDOM_LEN)
	// var ptest *byte = &test

	TestGeneratorLOG(test)
	fmt.Println("2:", test)
}

func TestGeneratorLOG(data []byte) {

	//ucdata := (*C.char)(unsafe.Pointer(&data[0]))
	//var data *[32]byte

	data[0] = 1

	// C.TestGeneratorLOG(cflag, ucdata, cslen)

	fmt.Println("1:", data)
	//	copy(data, ucdata)

}
