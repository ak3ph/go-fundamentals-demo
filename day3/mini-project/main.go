package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var students []Student

	for {
		var choice int
		fmt.Println("\n1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			var age int
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter age: ")
			fmt.Scanln(&age)

			students = append(students, Student{name, age})
			fmt.Println("Student added")

		case 2:
			fmt.Println("Student List:")
			for _, s := range students {
				fmt.Println(s.Name, "-", s.Age)
			}

		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
