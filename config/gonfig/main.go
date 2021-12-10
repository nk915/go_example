package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Logging struct {
	LogLevel string
	Test     string
}

type Config struct {
	Port    int `env:"APP_PORT"`
	Logging map[string]interface{}
}

func main() {

	config := Config{}
	err := gonfig.GetConf(getFileName(), &config)

	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	fmt.Println("Port: ", config.Port)
	fmt.Println("Logging: ", config.Logging)
	fmt.Println("Logging->LogLevel: ", config.Logging["LogLevel"])
	fmt.Println("Logging->LogLevel Type: ", reflect.TypeOf(config.Logging["LogLevel"]))
	fmt.Println("Logging->LogLevel String: ", config.Logging["LogLevel"].(string))
	//fmt.Println(config.Logging["Product"])
	//fmt.Println(config.Logging["NullPoint"])

}

func getFileName() string {
	env := os.Getenv("ENV")

	if len(env) == 0 {
		env = "dev"
	}

	filename := []string{"config/", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filepath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	fmt.Println("file_path: ", filepath)
	return filepath
}
