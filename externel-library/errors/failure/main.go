package main

import (
	"fmt"

	"github.com/morikuni/failure"
)

const (
	NotFound        failure.StringCode = "NotFound"
	InvalidArgument failure.StringCode = "InvalidArgument"
	Internal        failure.StringCode = "Internal"
)

func main() {
	err := Bar()
	fmt.Println("")
	fmt.Println("=====")
	fmt.Println(err)
	fmt.Printf("%+v\n", err.Error())
	fmt.Println("")
	fmt.Println("=====")
	fmt.Printf("%+v\n", err)
	fmt.Println("=====")
	//fmt.Printf("%+v\n", failure.Wrap(err, failure.WithCallStackSkip(1)))
	//	cs, _ := failure.CallStackOf(err)
	//	fmt.Printf("CallStack = %+v\n", cs)
	//
	//	fmt.Println("=====")
	//	fmt.Printf("Cause = %+v\n", failure.CauseOf(err))
	//
	//	fmt.Println("=====")
	//	ms, _ := failure.MessageOf(err)
	//	fmt.Printf("Message = %+v\n", ms)

}

func Bar() error {
	err := Foo("hello", "world")
	//err := Zerr()
	if err != nil {
		//return err
		return failure.Wrap(err, failure.Messagef("Message %s", "1"))
		// return failure.Translate(err, InvalidArgument)
		//		return failure.Wrap(err)
	}
	return nil
}

func Foo(a, b string) error {
	return failure.New(InvalidArgument,
		//failure.Context{"a": a, "b": b},
		failure.Message("Given parameters are invalid!!"),
	)
}
