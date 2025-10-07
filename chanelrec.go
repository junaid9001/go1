package main 
import(
	"go1/channelgo"
	"fmt"
	"time"
)


var helo=make(chan int)

func reciever(){
	for val:=range helo{
		fmt.Println(val)
	}
}

func main(){
go channelgo.Channel1(helo)
go reciever()
time.Sleep(1*time.Second)
}