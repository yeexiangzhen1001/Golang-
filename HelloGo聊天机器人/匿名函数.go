package main

import "fmt"

func proc (input string, processor func (str string)) {
	processor(input)
}
func main() {
	proc ("wnagxioaer", func(str string) {
		for _,v := range str {
			fmt.Printf("%c\n",v)
		}
	})
}
