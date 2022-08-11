package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type User struct {
	Name    string `json:"name,omitempty" xml:"name" param:"name"`
	Address string `json:"address" xml:"address" param:"address"`
	Num     int32  `json:"num" xml:"num" param:"num"`
	Test    string `json:"test,omitempty" xml:"test" param:"test"`
	Age     string `json:"age,omitempty"`
}

func main() {
	user := User{
		Name:    "nk915",
		Address: "Seoul",
		Num:     123,
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
		key, ok := ref.Type().Field(i).Tag.Lookup("json")
		value := fmt.Sprintf("%v", ref.Field(i))
		if ok && len(value) > 0 {
			//fmt.Printf("v(%s) --> %s : %s\n", v, ref.Type().Field(i).Tag.Get("json"), ref.Field(i))
			if len(strings.Split(key, ",")[0]) > 0 {
				fmt.Printf("%s : %s \n", strings.Split(key, ",")[0], value)
			}

		}
	}

	fmt.Println(len(params))
	baseUrl.RawQuery = params.Encode()
	fmt.Printf("URL : %s\n", baseUrl.String())

}
