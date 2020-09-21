package main

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Hello World")
}
