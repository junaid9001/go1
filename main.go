package main

import "fmt"

//check odd or even

func oddeven() {
	var num int
	fmt.Println("check even or odd")
	fmt.Println("enter number")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

//calculator

func calc() {
	var num1 int
	var typ string
	var num2 int
    fmt.Println("calculator")
	fmt.Println("enter first number")
	fmt.Scanln(&num1)
	fmt.Println("choose calculation ( +   -   *   /  )")
	fmt.Scanln(&typ)
	fmt.Println("choose second number")
	fmt.Scanln(&num2)

	switch typ {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	}

}

func main() {
	oddeven()
	calc()
}
