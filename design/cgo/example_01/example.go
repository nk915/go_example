package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lexample
#include "example.h"
*/
import "C"

func example() {
	C.exampleFunction()
}
