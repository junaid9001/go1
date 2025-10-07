package main
import (
	"fmt"
	"sync"
	"time"
)

var counter int

var mute sync.Mutex

func inc(){
mute.Lock()
counter++
mute.Unlock()
}


func main(){
	go inc()
	for i:=0;i<=9;i++{
		go inc()
	}
	time.Sleep(1*time.Second)
	fmt.Println(counter)
}