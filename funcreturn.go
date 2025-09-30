package main

import "fmt"


func function(num1, num2 int, fl1, fl2 float64) (sum, diff, mult int, divide float64) {
	sum = num1 + num2
	diff = num1 - num2
	mult = num1 * num2
	divide = fl1 / fl2
	return
}

func main() {
	sum, difference, product, qiotient := function(10, 20, 100, 10)
	fmt.Println("sum :", sum, "|diffreence:", difference, "|product:", product, "|qiotient", qiotient)

}
