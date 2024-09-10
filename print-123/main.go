package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)
	chan4 := make(chan int)

	go print(1, chan1, chan2)
	go print(2, chan2, chan3)
	go print(3, chan3, chan1)

	chan1 <- 1

	// wait forever
	<-chan4
}

func print(i int, input, output chan int) {
	for {
		<-input

		fmt.Println(i)

		output <- 1
	}
}
