package main
import "fmt"



func def(){
	
 fmt.Println("Start")

    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()

    panic("Something went wrong!")


}

func main(){
  def()
}