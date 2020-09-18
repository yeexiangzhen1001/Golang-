package main
import (
	"fmt"
	"reflect"
)
func main() {
	name := "xioaming"
	valueOfName := reflect.ValueOf(name)
	fmt.Println(valueOfName.Bytes())
}