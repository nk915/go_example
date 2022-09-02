package main

import (
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
)

type Person struct {
	Name  string `json:"name"`
	Phone string
}

var encoder = schema.NewEncoder()

func main() {

	baseUrl, _ := url.Parse("https://ncloud.apigw.ntruss.com/vpostgresql/v2")
	person := Person{"Jane Doe", "555-5555"}
	form := url.Values{}

	err := encoder.Encode(person, form)

	if err != nil {
		// Handle error
	}

	fmt.Printf("%v\n", form)
	baseUrl.RawQuery = form.Encode()
	fmt.Printf("%s\n", baseUrl.String())
	// Use form values, for example, with an http client
	//    client := new(http.Client)
	//    res, err := client.PostForm("http://my-api.test", form)
}
