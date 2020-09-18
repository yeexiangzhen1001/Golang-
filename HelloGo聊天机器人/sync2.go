package main
import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var rwLock sync.RWMutex
func main() {
	for i:=0;i<5;i++ {
		go func (i int ) {
			rwLock.RLock()
			defer rwLock.RUnlock()
			fmt.Println("read func " + strconv.Itoa(i)+" get wlock at" + time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second/10)
	for  i:=0;i<5;i++{
		go func (i int) {
			rwLock.Lock()
			defer rwLock.Unlock()
			fmt.Println("write func " + strconv.Itoa(i)+" get wlock at" + time.Now().String())
			time.Sleep(time.Second)

		}(i)
	}
	time.Sleep(time.Second*10)
}
