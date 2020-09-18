package main
import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)

	for i:=0;i<5;i++ {
		go func (i int) {
			fmt.Println("work" + strconv.Itoa(i) + " is done at "+ time.Now().String())
			time.Sleep(time.Second*2)
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	fmt.Println("all works done at "+time.Now().String())
}
