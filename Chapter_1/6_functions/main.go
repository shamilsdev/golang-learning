package main

import "fmt"

func add(a int, b int) {
	fmt.Println("add:", a+b)
}

func divide(a float64, b float64) float64 {
	return a / b
}

func multiply(x ...int) int {
	var result int = 1

	for _, n := range x {
		result *= n
	}

	return result
}

func add3(x, y, z int) (sum int) {
	sum = x + y + z
	return
}

func calc(x, y int) (sum, mul, div int) {
	sum = x + y
	mul = x - y
	div = x / y
	return
}

func main() {
	add(10, 5)

	fmt.Println("divide:", divide(10, 3))

	fmt.Println(add3(4, 4, 5))

	fmt.Println(calc(100, 5))
}
