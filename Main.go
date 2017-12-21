package main

import (
	"fmt"
)

func main() {
	// var x bool
	// x = false

	x := false

	if x {
		fmt.Println("X is true")
	} else {
		fmt.Println("X is false")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	var z int
	z = 100
	for {
		z++
		fmt.Println(z)

		if z > 104 {
			break
		}
	}

	cont := true
	for cont {
		fmt.Println("continue")
		cont = false
	}

	numbers := []int{9, 8, 7, 6, 5}
	for index, val := range numbers {
		fmt.Printf("index=%v value=%v\n", index, val)
	}
}
