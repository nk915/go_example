package example

/*
#cgo CFLAGS: -I./example/h
#cgo LDFLAGS: -L./example/lib -lexample
#include "example.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type ScanLib struct {
	Init  unsafe.Pointer
	Clean C.fnVCScan_Clean
}

func example() {
	// C 함수 호출
	C.exampleFunction()

	// C 구조체 초기화
	var es C.ExampleStruct
	es.id = 1

	// 문자열 복사 (Go 문자열 -> C 문자열)
	name := "Golang User"
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.strcpy(&es.name[0], cName)

	// C 함수 호출 (구조체를 인자로 전달)
	C.printExampleStruct(es)

	scan := ScanLib{}
	fmt.Println(scan)
	fmt.Println("c.max:", C.MAX)
}
