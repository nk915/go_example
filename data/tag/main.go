package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type User struct {
	Name    string `json:"name,omitempty" xml:"name" param:"name"`
	Address string `json:"" xml:"address" param:"address"`
	Test    string `json:"test,omitempty" xml:"test" param:"test"`
	Age     string `json:"age,omitempty"`
}

type UserA struct {
	Name    *string `json:"name,omitempty" xml:"name" param:"name"`
	Address *string `json:"address,omitempty" xml:"address" param:"address"`
	Test    *string `json:"test,omitempty" xml:"test" param:"test"`
	Age     *string `json:"age,omitempty"`
}

func main() {
	//	t := reflect.TypeOf(User{})
	//	for i := 0; i < t.NumField(); i++ {
	//		v, ok := t.Field(i).Tag.Lookup("json")
	//		if ok {
	//			fmt.Printf("(%14s) is Json Field : %s\n", t.Field(i).Name, v)
	//		}
	//		v, ok = t.Field(i).Tag.Lookup("xml")
	//		if ok {
	//			fmt.Printf("(%14s) is xml Field : %s\n", t.Field(i).Name, v)
	//		}
	//	}

	user := User{
		Name:    "nk915",
		Address: "Seoul",
		Age:     "",
		Test:    "",
	}

	//ref := reflect.TypeOf(user)
	ref := reflect.ValueOf(user)
	fmt.Println(ref.NumField())
	fmt.Println(ref.IsValid())

	baseUrl, _ := url.Parse("https://ncloud.apigw.ntruss.com/vpostgresql/v2")
	//baseUrl.Path += "addCloudPostgresqlDatabaseList"
	baseUrl.Path += ""

	params := url.Values{}
	//params.Add("regionCode", "KR")
	//params.Add("cloudPostgresqlInstanceNo", "****890")

	/*
	   GET {API_URL}/addCloudPostgresqlDatabaseList
	   ?regionCode=KR
	   &cloudPostgresqlInstanceNo=****890
	   &cloudPostgresqlDatabaseList.1.name=pgtest
	   &cloudPostgresqlDatabaseList.1.owner=testuser
	*/

	for i := 0; i < ref.NumField(); i++ {
		//	v, ok := ref.Type().Field(i).Tag.Lookup("param")
		//	if ok && len(ref.Field(i).String()) > 0 {
		//		fmt.Printf("v(%s) --> %s : %s\n", v, ref.Type().Field(i).Tag.Get("param"), ref.Field(i))
		//		params.Add(v, ref.Field(i).String())
		//		//fmt.Printf("(%14s) is Json Field : %s \n", ref.Field(i).Name, v)
		//	}

		v, ok := ref.Type().Field(i).Tag.Lookup("json")
		if ok && len(ref.Field(i).String()) > 0 {
			//fmt.Printf("v(%s) --> %s : %s\n", v, ref.Type().Field(i).Tag.Get("json"), ref.Field(i))

			if len(strings.Split(v, ",")[0]) > 0 {
				fmt.Printf("%s : %s \n", strings.Split(v, ",")[0], ref.Field(i))
			}

		}

		//		v, ok = ref.Field(i).Tag.Lookup("xml")
		//		if ok {
		//			fmt.Printf("(%14s) is xml Field : %s \n", ref.Field(i).Name, v)
		//		}
	}

	fmt.Println(len(params))
	baseUrl.RawQuery = params.Encode()
	fmt.Printf("URL : %s\n", baseUrl.String())

}
