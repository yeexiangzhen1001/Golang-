package main

import (
	"fmt"
	"time"
)

func main(){
	go func (name string) {
		fmt.Println("Hello"+name)

	}("xuan")
	time.Sleep(time.Second)
}
