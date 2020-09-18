package main

import (
	"fmt"
	"time"
)

func consume(ch chan int) {
	time.Sleep(time.Second*100)
	<-ch
}

func main() {
	ch := make(chan int,2)
	go  consume(ch)

	ch <- 0
	ch <- 1
	fmt.Println("I am free!")

	time.Sleep(time.Second)
}
