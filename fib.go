package main

import "fmt"

func fibonacci() {
	frst, scnd := 0, 1

	for i := 0; i <= 10; i++ {
		fmt.Println(frst)
		frst, scnd = scnd, frst+scnd
	}
}

func main() {
	fibonacci()
}
