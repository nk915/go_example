package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lhello
#include "hello.c"
*/
import "C"

func main() {
	// let's call it
	C.hello()

}
