package main

import (
	"fmt"
)

type Animal interface {
	call()
}

type Cat struct {
}

type Dog struct {
}

func (dog *Dog) call() {
	fmt.Println("Bark")
}

func (cat *Cat) call() {
	fmt.Println("Meow")
}

func main() {
	var animal Animal

	animal = &Dog{}
	animal.call()

	animal = &Cat{}
	animal.call()
}
