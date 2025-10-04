package main

import "fmt"

func maap() {
	students := map[string]int{
		"Abi":   45,
		"Alice": 40,
		"Mark":  63,
		"Zoe":   48,
	}

	names := []string{
		"Alice",
		"Bob",
		"Charlie",
		"Diana",
		"Ethan",
		"Fiona",
		"George",
		"Hannah",
		"Ian",
		"Julia",
	}

	numbers := []int{12, 25, 37, 44, 59, 63, 71, 85, 92, 100}

	fmt.Println(students,"\n")

	for key, value := range students {
		fmt.Printf("key :%v ,value: %v |", key, value)
	}

	students["junaid"] = 50
	students["abi"] = 50
	delete(students, "alice")

	fmt.Println(students,"\n")

	for i := 0; i <= 9; i++ {
		students[names[i]] = numbers[i]
	}

	fmt.Println(students,"\n")

   for key, value:= range students {
      if key=="Julia"{
         delete(students,key)
      }

      if value>80{
         delete(students,key)
      }
	}

   fmt.Println(students)

}

func main() {
	maap()
}

















