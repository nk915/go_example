package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	examSliceChange()
	examIndexFunc()
}

func examSliceChange() {
	type InnerConfig struct {
		A string
		B string
	}
	type Config struct {
		Inner []InnerConfig
	}

	in1 := []InnerConfig{
		InnerConfig{A: "1", B: "2"},
		InnerConfig{A: "3", B: "4"},
		InnerConfig{A: "5", B: "6"},
	}
	in2 := []InnerConfig{
		InnerConfig{A: "A", B: "B"},
	}

	config := Config{}
	fmt.Printf("config empty: %+v\n", config)

	config.Inner = in1
	fmt.Printf("config insert in1: %+v\n", config)

	config.Inner = in2
	fmt.Printf("config insert in2: %+v\n", config)

}

func examIndexFunc() {
	type Config struct {
		Key   string
		Value string
	}

	// I form a slice of the above struct
	myconfig := []Config{
		Config{Key: "1", Value: "A"},
		Config{Key: "2", Value: "B"},
		Config{Key: "2", Value: "C"},
	}

	idx := slices.IndexFunc(myconfig, func(c Config) bool { return c.Key == "2" })

	fmt.Println(idx)

}
