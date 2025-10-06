package main

import (
	"fmt"
	"time"
)

func odds() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("from odd function ", i)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func even() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("from even function ", i)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func main() {
	go odds()
	go even()

	time.Sleep(1 * time.Second)
}
