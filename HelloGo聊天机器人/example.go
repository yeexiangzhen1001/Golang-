package main

import "fmt"

type Swimming struct {

}
func (swim *Swimming) swim() {
	fmt.Printf("swimming is my ability\n" +
		"")
}

type Flying struct {

}

func (fly *Flying) fly() {
	fmt.Printf("flying is my ability\n")
}

type WildDuck struct{
	Swimming
	Flying
}

type DomesticDuck struct {
	Swimming
}

func main() {
	wild := WildDuck{}
	wild.fly()
	wild.swim()

	domestic := DomesticDuck{}
	domestic.swim()
}
