package main

import "fmt"

func add(a int, b int) {
	fmt.Println("add:", a+b)
}

func divide(a int, b int) int {
	return a / b
}

func multiply(x ...int) int {
	var result int = 1

	for _, n := range x {
		result *= n
	}

	return result
}

func main() {
	add(10, 5)

	fmt.Println("divide:", divide(10, 2))
}
