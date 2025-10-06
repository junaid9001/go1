package main 

import "fmt"


	type Person struct{
		name string
		age int
	}

	pslice:=Person[]{

	}


    func(reciever Person)Greet(){
		fmt.Printf("your name is %v, and your age is %v",reciever.name,reciever.age)
	}


func main(){
	p:=Person{name:"juanid",age:21,}
	p.Greet()
}