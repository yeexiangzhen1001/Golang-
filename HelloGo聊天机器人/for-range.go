package main

import (
	"fmt"
)

func main() {
	nums := [...]int {1,2,3,4,5,6,7,8,9,10}
	for k,v := range nums{
		fmt.Println(k,v," ")
	}
	fmt.Println()

	slis := []int{1,2,3,4,5,6,7,8,9,10}
	for k, v := range slis {
		fmt.Println(k,v," ")
	}
	fmt.Println()

	tmpMap := map[int]string {
		0 : "小明",
		1 : "xiaohong",
		2 : "xaiobai",
	}
	for k, v := range tmpMap {
		fmt.Println(k,v," ")
	}
}
