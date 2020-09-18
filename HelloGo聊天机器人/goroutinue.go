package main

import (
	"fmt"
	_"time"
)

func test() {
	fmt.Println("I am work in a single goroutinue")
}

func main() {
	go test()

	//time.Sleep(time.Second)
}
