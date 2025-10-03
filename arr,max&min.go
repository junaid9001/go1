package main
import "fmt"
func handle(){
	var arr=[...]int{10,100,250,4,75,1000}
	var min =0
	var max =0
	for _,value:=range arr{
       if value>max{
		max=value
	   }
	   if value<min{
         min=value
	   }
	}
	fmt.Printf("max num is %v , min num is %v",max,min)
}

func main(){
handle()
}