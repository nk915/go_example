package main

type User struct {
	Name         string
	Role         string
	Age          int32
	EmployeeCode int64 `copier:"EmployeeNum"` // specify field name
	Test         string

	// Explicitly ignored in the destination struct.
	Salary int
}

type User2 struct {
	Name         *string
	Role         *string
	Age          int32
	EmployeeCode *int64 `copier:"EmployeeNum"` // specify field name
	Test         string

	// Explicitly ignored in the destination struct.
	Salary int
}

func main() {
	test_1()
}
