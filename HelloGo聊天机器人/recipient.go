package main
import "fmt"

type Person struct {
	Name string
	Brith string
	ID  int64
}

func (person *Person) changName (name string) {
	person.Name = name
}

func (person Person) printMess()  {
	fmt.Printf("My name is %v,and my brithday is %v,and my id is %v\n",person.Name,person.Brith,person.ID)

	person.ID = 3
}

func main() {
	p1 := Person{
		Name:  "xiaowang",
		Brith: "1998-10-1",
		ID:    1,
	}
	p1.printMess()
	p1.changName("laowang")
	p1.printMess()
}

