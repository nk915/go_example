package main

import "fmt"

func main() {
	ch := make(chan string, 10)
	defer close(ch)

	sendChan(ch)
	//	receiveChan(ch)

	loopChan(ch)

}

func sendChan(ch chan<- string) {
	ch <- "Data"
	ch <- "Data"
	// x := <-ch // 에러발생
}

func loopChan(ch chan string) {
	fmt.Println("start loop chan..")

	for len(ch) > 0 {
		ch <- "---"
		fmt.Println(<-ch)
		fmt.Println(<-ch)

		//	if i, success := <-ch; success {
		//		println(i)
		//	} else {
		//		println("else ", i)
		//		break
		//	}
	}

	fmt.Println("end loop chan..")
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
