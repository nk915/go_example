package main

import (
	"fmt"
	"net/url"
	"reflect"
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

	UrlTest(user)
	FieldChange(user)
	user.ToFrom([]string{"Name", "Num", "Address", "AA"})
}

func (user *User) ToFrom(list []string) error {
	copy_user := User{}
	fmt.Println(user, list)

	// pointer to struct - addressable
	org := reflect.ValueOf(user).Elem()
	new := reflect.ValueOf(&copy_user).Elem()

	if org.Kind() != new.Kind() {
		fmt.Println("False")
		return nil
	}

	// New 구조체가 삽입 가능한지 확인한다.
	for _, field := range list {
		orgf := org.FieldByName(field)
		newf := new.FieldByName(field)

		if newf.Kind() == reflect.String && orgf.Kind() == reflect.String {
			newf.SetString(orgf.String())
		}
	}

	fmt.Println(copy_user, list)
	return nil
}

func FieldChange(user User) {

	fmt.Println("Before: ", user)
	// pointer to struct - addressable
	ps := reflect.ValueOf(&user)
	// struct
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		// exported field
		f := s.FieldByName("Address")
		if f.IsValid() {
			fmt.Println("TEST")
			// A Value can be changed only if it is addressable and was not obtained by the use of unexported struct fields.
			if f.CanSet() {
				fmt.Println("TEST")
				// change value of N
				if f.Kind() == reflect.Int32 {
					x := int64(7)
					fmt.Println("TEST")
					if !f.OverflowInt(x) {
						fmt.Println("TEST")
						f.SetInt(x)
					}
				}

				if f.Kind() == reflect.String {
					x := string("TEST")
					f.SetString(x)
				}
			}
		}
	}

	fmt.Println("After: ", user)
}

func UrlTest(user User) {
	ref := reflect.ValueOf(user)
	fmt.Println(ref.NumField())
	fmt.Println(ref.IsValid())

	baseUrl, _ := url.Parse("https://ncloud.apigw.ntruss.com/vpostgresql/v2")
	//baseUrl.Path += "addCloudPostgresqlDatabaseList"
	baseUrl.Path += ""

	params := url.Values{}
	//params.Add("regionCode", "KR")
	//params.Add("cloudPostgresqlInstanceNo", "****890")

	for i := 0; i < ref.NumField(); i++ {
		key, ok := ref.Type().Field(i).Tag.Lookup("json")
		value := fmt.Sprintf("%v", ref.Field(i))
		if ok && len(value) > 0 {
			fmt.Printf("v(%s) --> %s : %s\n", key, ref.Type().Field(i).Tag.Get("json"), ref.Field(i))
			//	if len(strings.Split(key, ",")[0]) > 0 {
			//		fmt.Printf("%s : %s \n", strings.Split(key, ",")[0], value)
			//	}

		}
	}
	fmt.Println(len(params))
	baseUrl.RawQuery = params.Encode()

	/*
	   GET {API_URL}/addCloudPostgresqlDatabaseList
	   ?regionCode=KR
	   &cloudPostgresqlInstanceNo=****890
	   &cloudPostgresqlDatabaseList.1.name=pgtest
	   &cloudPostgresqlDatabaseList.1.owner=testuser
	*/
	fmt.Printf("URL : %s\n", baseUrl.String())
}
