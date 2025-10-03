package main

import "fmt"

func nodupe() {
	slic := []int{11, 11,2,5, 12, 15, 5, 4, 1,1,1,1,8, 6, 5, 4, 7, 2, 1, 6, 8, 2}
	var emslic = []int{}
	fmt.Println("before ", slic)

	for i := 0; i < len(slic); i++ {
		var boo bool = true
		for j := i + 1; j < len(slic); j++ {
			if slic[i] == slic[j] {
				boo = false
				break
			}
		}
		if boo{
			emslic = append(emslic, slic[i])
			
		}
	}
	slic=emslic
	fmt.Println("after ", slic)
}

func main() {
	nodupe()
}
