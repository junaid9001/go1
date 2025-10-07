package main

import "fmt"

type Student struct {
	name string
	age  int
}

var students = []Student{
	{name: "Alice", age: 21},
	{name: "Joe", age: 19},
}

func Addstudent(nam string, ag int) {
	students = append(students, Student{name: nam, age: ag})
}

func Liststudents() {
	for index, value := range students {
		fmt.Printf(" no:%v  name:%v  age:%v\n", index, value.name, value.age)
	}
}

func Update() {
	var index int
	var typee int
	var name string
	var age int
	fmt.Println("enter index of student")
	fmt.Scanln(&index)
	fmt.Println("1.edit name\n 2.edit age\n 3.edit age and name")
	fmt.Scanln(&typee)
	switch typee {
	case 1:
		fmt.Println("enter name")
		fmt.Scanln(&name)
		students[index].name = name
	case 2:
		fmt.Println("enter age")
		fmt.Scanln(&age)
		students[index].age = age

	case 3:
		fmt.Println("enter name")
		fmt.Scanln(&name)
		fmt.Println("enter age")
		fmt.Scanln(&age)
		students[index].name = name
		students[index].age = age
	}

}

func Deletestudent(index int) {
	students = append(students[:index], students[index+1:]...)
}

func main() {
	for {
		var choice int
		fmt.Println("\n 1.Create new students\n 2.Read list of students\n 3.Update a students name or age\n 4.Delete a student")
		fmt.Println("Enter Your Choice")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var name string
			var age int
			fmt.Println("Enter name")
			fmt.Scanln(&name)
			fmt.Println("Enter age")
			fmt.Scanln(&age)
			Addstudent(name, age)

		case 2:
			Liststudents()

		case 3:
			Update()

		case 4:
			var index int
			fmt.Println("Enter index number of student to remove")
			fmt.Scanln(&index)
			Deletestudent(index)

		case 5:
			return

		default:
		}

	}
}
