package main

import (
	"fmt"
	"time"

	"github.com/nathan-osman/go-sunrise"
)

func main() {
	rise, set := sunrise.SunriseSunset(
		43.65, -79.38, // Toronto, CA
		2000, time.January, 1, // 2000-01-01
	)

	fmt.Printf("sunrise: %v\nsunset: %v\n", rise, set)
}
