package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type aa struct {
	a string
	b bb
}

type bb struct {
	b string
}

func main() {

	// using standard library "flag" package
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // retrieve value from viper

	fmt.Printf("i : %d\n", i)

	viper.Set("aa", 1)
	fmt.Println(viper.Get("aa"))
	viper.Set("aa", 2)
	fmt.Println(viper.Get("aa"))

	viper.Set("aa", aa{a: "a", b: bb{b: "b"}})
	fmt.Println(viper.Get("aa"))

}
