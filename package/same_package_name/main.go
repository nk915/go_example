package main

import (
	"fmt"

	custom_lib "same_package_name_test.com/nk915/lib"
)

func main() {
	custom_lib.PrintLib1()
	custom_lib.PrintLib2()

	fmt.Println()

	custom_lib.PrintLibBy1()
	custom_lib.PrintLibBy2()
}
