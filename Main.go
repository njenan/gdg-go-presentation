package main

import (
	"fmt"
)

func toDefer(i int) {
	fmt.Printf("Defer func execute count = %v\n", i)
}

func badFunc() {
	panic("Something went wrong")
}

func main() {
	defer toDefer(1)
	defer toDefer(2)
	defer toDefer(3)
	//defer func() {
	//	fmt.Printf("Recovered from '%v'\n", recover())
	//}()

	fmt.Println("Main function body")

	badFunc()
	recover()

	fmt.Println("After bad function")
}
