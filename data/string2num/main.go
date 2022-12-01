package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString("abc123def987asdf", -1))
	fmt.Println(re.FindAllString("eth0", -1))

	fmt.Println(re.FindAllString("-1212awsd12123", -1))
}
