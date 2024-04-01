package main

import (
	"flag"
	"fmt"
)

func main() {
	var Name string
	name := flag.String("host_name", "aa", "hostname")
	isServer := flag.Bool("server", false, "Run as server")
	isClient := flag.Bool("client", false, "Run as client")
	flag.Parse()

	Name = *name
	fmt.Println(Name)
	if *isServer {
		fmt.Println("Running as server...")
		// 서버 관련 코드
	} else if *isClient {
		fmt.Println("Running as client...")
		// 클라이언트 관련 코드
	} else {
		fmt.Println("Please specify -server or -client")
	}
}
