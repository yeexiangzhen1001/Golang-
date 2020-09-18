package main
import (
	"flag"
	"fmt"
)

func main () {
	surname := flag.String("surname","王","姓")
	var personalName string
	flag.StringVar(&personalName,"personalName","小二","名")
	id := flag.Int("id",0,"ID")
	flag.Parse()
	fmt.Printf("I am %v %v,and my id is %v\n",*surname, personalName ,*id)
}
