package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name         string
	Role         string
	Age          int32
	EmployeeCode int64 `copier:"EmployeeNum"` // specify field name
	Test         string
	SystemId     string

	Index int32
	// Explicitly ignored in the destination struct.
	Salary int
}

type User2 struct {
	Name         *string
	Role         *string
	Age          int32
	EmployeeCode *int64 `copier:"EmployeeNum"` // specify field name
	Test         string
	SystemType   string `copier:"SystemId"`
	Index        int16
	// Explicitly ignored in the destination struct.
	Salary int
}

func main() {
	//	test_1()
	test_3()
}

func test_3() {
	from := User2{SystemType: "AA", Index: 123}
	to := User{}

	copier.Copy(&to, from)
	fmt.Println(to)
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

// Tags in the destination Struct provide instructions to copier.Copy to ignore
// or enforce copying and to panic or return an error if a field was not copied.
type Employee struct {
	// Tell copier.Copy to panic if this field is not copied.
	Name string `copier:"must"`

	// Tell copier.Copy to return an error if this field is not copied.
	Age int32 `copier:"must,nopanic"`

	// Tell copier.Copy to explicitly ignore copying this field.
	Salary int `copier:"-"`

	Test       string
	DoubleAge  int32
	EmployeeId int64 `copier:"EmployeeNum"` // specify field name
	SuperRole  string
}

func (employee *Employee) Role(role string) {
	employee.SuperRole = "Super " + role
}

func test_1() {
	from := User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000, Test: "TEST"}
	to := User2{}

	fmt.Println(from)
	fmt.Println(to)

	copier.CopyWithOption(&to, &from, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	fmt.Println(to)

}

func test_2() {
	var (
		user      = User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000, Test: "TEST"}
		users     = []User{{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 100000}, {Name: "jinzhu 2", Age: 30, Role: "Dev", Salary: 60000}}
		employee  = Employee{Salary: 150000}
		employees = []Employee{}
	)

	copier.Copy(&employee, &user)

	fmt.Printf("%#v \n", employee)
	// Employee{
	//    Name: "Jinzhu",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    Salary:150000,            // Copying explicitly ignored
	//    DoubleAge: 36,            // Copy from method
	//    EmployeeId: 0,            // Ignored
	//    SuperRole: "Super Admin", // Copy to method
	// }

	// Copy struct to slice
	copier.Copy(&employees, &user)

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeeId: 0, SuperRole: "Super Admin"}
	// }

	// Copy slice to slice
	employees = []Employee{}
	copier.Copy(&employees, &users)

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeeId: 0, SuperRole: "Super Admin"},
	//   {Name: "jinzhu 2", Age: 30, Salary:0, DoubleAge: 60, EmployeeId: 0, SuperRole: "Super Dev"},
	// }

	// Copy map to map
	map1 := map[int]int{3: 6, 4: 8}
	map2 := map[int32]int8{}
	copier.Copy(&map2, map1)

	fmt.Printf("%#v \n", map2)
	// map[int32]int8{3:6, 4:8}
}
