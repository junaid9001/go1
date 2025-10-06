package main
import "fmt"


type Multiply interface{
	multiply() int
}

type Person struct{
	mult int
}

func (p Person) multiply()int{
	return 1010
}

func multyplyy(m Multiply){
	fmt.Println()
}


func main (){
  p:=Person{
	mult:100,
  }
  multyplyy(p)
}
