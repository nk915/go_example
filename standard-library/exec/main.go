package main

import (
	"fmt"
	"os/exec"
)

func main() {
	if err := exec.Command("chronyc", "manual", "on").Run(); err != nil {
		panic(err)
	}

	if output, err := exec.Command("chronyc", "settime", "2024-03-28 17:00:00").CombinedOutput(); err != nil {
		panic(err)

	} else {
		fmt.Println(output)
	}

}
