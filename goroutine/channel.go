package main

import (
	"fmt"
	"time"
)

var num = make(chan int)

func Producer() {
	num <- 10
	fmt.Println("producer")
}

func Consumer() {
	rec := <-num
	fmt.Println(rec)
}

func main() {
	go Producer()
	go Consumer()

	time.Sleep(1 * time.Second)
}
