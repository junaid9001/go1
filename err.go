package main

import (
	"errors"
	"fmt"
)

func handlerror(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("cant do this dawg")
	}
	return num1 / num2, nil
}

func main() {
	divi, err := handlerror(10, 0)
	if err != nil {
		fmt.Println(divi, err)
	} else {
		fmt.Println(divi)
	}

}
