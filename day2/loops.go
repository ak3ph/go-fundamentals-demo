package main

import "fmt"

func loops() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("")
	}
}

func whileLoops() {
	i := 1
	for i <= 5 {
		fmt.Printf("i")
	}
}

func infiniteLoops() {
	for {
		fmt.Println("This will run forever!")
	}
}
