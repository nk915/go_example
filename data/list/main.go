package main

import "fmt"

func main() {
	listPlist()
	fmt.Println("-----------------")
	listGet()
	fmt.Println("-----------------")

	listS()
	fmt.Println("-----------------")
	inner()

	fmt.Println("-----------------")
	innermap()
}

type key struct {
	K string
}

type test struct {
	A     string
	B     string
	Inner InnerTest
}
type InnerTest struct {
	C string
}

func (in *InnerTest) update() {
	in.C = "C-update"
}

func (t *test) update() {
	t.A = "A-update"
}

func innermap() {
	temp := make(map[key]test)

	temp[key{K: "2"}] = test{A: "A"}
	temp[key{K: "1"}] = test{A: "B"}

	fmt.Printf("%+v\n", temp[key{K: "1"}])
	for _, v := range temp {

		fmt.Printf("1: %+v\n", v)
		v.A = "AAA"
		fmt.Printf("2: %+v\n", v)
	}
	fmt.Printf("%+v\n", temp[key{K: "1"}])

}

func inner() {
	tmp := test{
		A:     "A",
		B:     "B",
		Inner: InnerTest{C: "C"},
	}

	testlist := []test{}
	testlist = append(testlist, tmp, tmp)

	fmt.Printf("%+v\n", testlist)

	for i := range testlist {
		testlist[i].update()
	}
	fmt.Printf("%+v\n", testlist)
}

func listS() {
	testList := []test{}
	test := test{A: "A", B: "A"}

	testList = append(testList, test, test, test, test)

	fmt.Printf("%v\n", testList)

	for i, v := range testList {
		testList[i].B = "CC"
		v.A = "CCCC"
	}

	fmt.Println(testList)

}

func listPlist() {
	testList := []test{}

	fmt.Printf("(%d) %+v\n", len(testList), testList)

	testListA := []test{}
	testA := test{A: "A", B: "A"}
	testListA = append(testListA, testA)
	fmt.Printf("testListA: %+v\n", testListA)

	testList = append(testList, testListA...)
	fmt.Printf("testList: %+v\n", testList)

	testListB := []test{}
	testB := test{A: "B", B: "B"}
	testListB = append(testListB, testB)
	fmt.Printf("testListB: %+v\n", testListB)

	testList = append(testList, testListB...)
	fmt.Printf("testList: len(%d) %+v\n", len(testList), testList)
}

func listGet() {
	testListA := []test{}
	testA := test{A: "A", B: "A"}
	testListA = append(testListA, testA)

	test := &testListA[0]
	fmt.Printf("%+v\n", test)

	test.A = "TEST"
	fmt.Printf("%+v\n", test)
	fmt.Printf("%+v\n", testListA)

}
