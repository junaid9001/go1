package main
import "fmt"



func changeint(num,mult *int){
    *num=*num+100
	*mult=*mult*500
}


func main(){
  num:=100
  mult:=10
  changeint(&num,&mult)
  fmt.Println(num,mult)
}