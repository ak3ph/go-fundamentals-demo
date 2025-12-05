package main

import "fmt"

func slices() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)

	fmt.Println(numbers)
}
