package main

import (
	"fmt"
	"strconv"
	"sync"
)
var syncMap sync.Map
var waitGroup sync.WaitGroup
func main() {
	routineSize := 5
	waitGroup.Add(routineSize)
	for i := 0; i<routineSize; i++ {
		go addNumber(i *10)
	}
	waitGroup.Wait()
	var size int
	syncMap.Range(func(key,value interface{})bool {
		size++
		return true
	})
	fmt.Println("sync.Map current size is " +strconv.Itoa(size))
	value , ok := syncMap.Load(0);if ok{
		fmt.Println("key 0 has value",value," ")
	}
}
func addNumber(begin int){
	for i:=begin;i<begin;i++ {

		syncMap.Store(i,i)
	}
	waitGroup.Done()
}