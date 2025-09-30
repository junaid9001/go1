package main

import "fmt"

//check odd or even

func oddeven() {
	var num int
	fmt.Println("enter number")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

func main() {
	oddeven()
}
