package main

import (
	"fmt"
	"strconv"
	"errors"
)

func main() {
	/* Hello world */
	fmt.Println("Hello world")



	/* Variable declaration */
	var a int
	a = 1

	b := 2 // <- Can't use this outside of functions

	fmt.Println(a,b)


	var people [3]string
	people[0] = "Mike"
	people[1] = "John"
	people[2] = "Jane"



	var ages map[string]string
	ages = make(map[string]string)

	ages["John"] = "21"
	ages["Jane"] = "22"
	ages["Joe"] = "17"


	fmt.Println("John is " + ages["John"] + " years old")
	fmt.Println("Jane is " + ages["Jane"] + " years old")
	fmt.Println("Joe is " + ages["Joe"] + " years old")



	/* Functions */
	out, _ := favoriteNumber(7)

	fmt.Println(out)



	/* IF-ELSE statements */
	if (1 == 2) {
		fmt.Println("Impossible!")
	} else {
		fmt.Println("Makes sense")
	}


	/* For loops */
	for i := 0; i < 100; i ++ {
		fmt.Println(i)
	}

	/* Forever loops */
	c := 0
	for {
		if c == 2 {
			fmt.Println("Get me out of here!")
			break
		} else {
			fmt.Println(c)
		}
		c ++
	}

}


func favoriteNumber(i int) (string, error) {
	if (i == 9) {
		return "", errors.New("Nobody likes the number 9")
	} else {
		return strconv.Itoa(i) + " is my favorite number", nil
	}
}