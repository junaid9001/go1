package main 
import ("fmt"
"time")

func main(){
	name:=make(chan string)

	go func(){
		name <- "juanid"
	}()


	go func(){
		name <- "jamshad"
	}()
 

	 frst:=<-name
	 second:=<-name

    fmt.Println(frst,second)
	  
	time.Sleep(2*time.Second)
	
}