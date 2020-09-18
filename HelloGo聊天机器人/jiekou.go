package main

import (
	"fmt"
)

type Cat interface {

	CatchMouse ()
}

type Dog interface {

	Bark ()
}

type catDog struct {
	Name string
}

func (catDog *catDog) CatchMouse() {
	fmt.Printf("%v caught the mouse and ate it!\n",catDog.Name)
}

func (catDog *catDog) Bark() {
	fmt.Printf("%v barked loudly!\n",catDog.Name)
}

func main() {
	catDog := &catDog{
		"Lucy",
	}

	var cat Cat
	cat = catDog
	cat.CatchMouse()

	var dog Dog
	dog = catDog
	dog.Bark()

}
