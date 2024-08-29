package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("[service] Start...")

	for i := 10; i > 0; i-- {
		fmt.Printf("[service] wait %d...\n", i)
		time.Sleep(time.Second)
	}

	fmt.Println("[service] End...")
}
