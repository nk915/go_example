package main

import (
	"fmt"

	"local-testing.com/nk915/hello/country/en"
	"local-testing.com/nk915/hello/country/ko"
)

func main() {
	fmt.Println(ko.Hello())
	fmt.Println(en.Hello())
}
