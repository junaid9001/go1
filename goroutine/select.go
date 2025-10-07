package main
import ("fmt";"time")



var ch1=make(chan string,1)

var ch2=make(chan string,1)

func func1(){
	select{
	case ch1<-"channel 1":
		fmt.Println("ch1 sended")
	case ch2<-"channel 2":
		fmt.Println("ch2 sended")
	}
}

func func2(){
 str2:=<-ch2
 str1:=<-ch1
 fmt.Println(str1,str2)
 
}

func main(){
go func1()
go func2()
time.Sleep(1*time.Second)
}