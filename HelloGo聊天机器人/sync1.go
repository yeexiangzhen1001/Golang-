package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	var lock sync.Mutex
	go func () {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("func1 get lock at" + time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("func1 release lock" + time.Now().String())
	}()
	time.Sleep(time.Second/10)
	go func () {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("func2 get lock"+time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("func2 release lock"+time.Now().String())

	}()
	time.Sleep(time.Second*4)
}