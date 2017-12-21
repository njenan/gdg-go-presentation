package main

import (
	"fmt"
	"sync"
	"time"
)

func print(i int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println(i)
	wg.Done()
}

func calcSquare(i int, ch chan int) {
	ch <- i * i
}

func main() {
	/*
		var wg sync.WaitGroup
		wg.Add(3)

		go print(1, &wg)
		go print(2, &wg)
		go print(3, &wg)

		wg.Wait()

		fmt.Println("Done")
	*/

	channel := make(chan int)

	go calcSquare(9, channel)
	go calcSquare(15, channel)
	go calcSquare(31, channel)
	x, y, z := <-channel, <-channel, <-channel

	fmt.Printf("%v, %v, %v\n", x, y, z)

}
