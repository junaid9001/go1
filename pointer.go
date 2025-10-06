package main
import ("fmt") 


func pointer(a *int){
	*a=*a+10
}


func run(){
	fmt.Println("hello")
}

var mychan =make(chan int)

func main(){
	num:=100
	pointer(&num)
	fmt.Println(num)
	go run()
}