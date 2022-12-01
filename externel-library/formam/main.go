package main

import (
	"net/http"
	"net/url"
	"time"

	formam "github.com/monoculum/formam/v3"
)

type InterfaceStruct struct {
	ID   int
	Name string
}

type Company struct {
	Public     bool      `formam:"public"`
	Website    url.URL   `formam:"website"`
	Foundation time.Time `formam:"foundation"`
	Name       string
	Location   struct {
		Country string
		City    string
	}
	Products []struct {
		Name string
		Type string
	}
	Founders  []string
	Employees int64

	Interface interface{}
}

func MyHandler(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()

	m := Company{
		// it's is possible to access to the fields although it's an interface field!
		Interface: &InterfaceStruct{},
	}
	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "formam"})
	return dec.Decode(r.Form, &m)
}
