package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, err := strconv.Atoi("123a")

	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}

	fmt.Println(num)
}
