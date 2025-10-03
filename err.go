package main

import (
	"fmt"
)

func handlerror(num1, num2 int) (int,error) {
	if num2 == 0 {
		return 0,fmt.Errorf("cant divide %v by %v",num1,num2)
	}
	return num1 / num2,nil
	

}


func main() {
   sun,err:=handlerror(10,0)
  if err!=nil{
	fmt.Println(err)
  }else{
	fmt.Println(sun)
  }

}
