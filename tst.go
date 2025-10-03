package main
import "fmt"

func test(num int ,st string)(low int,str string){
low=num*num
str="hello"+st
return 
}

func main(){
	mult,hai:=test(10, " junaid ")
	fmt.Println(mult,hai)
}