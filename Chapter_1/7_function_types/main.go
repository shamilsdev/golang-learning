package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func action(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func choose(m int) func(int, int) int {
	switch m {
	case 1:
		return add
	case 2:
		return minus
	default:
		return add
	}
}

func main() {

	var f func(int, int) int
	f = minus
	fmt.Println(f(3, 2))

	v := action(100, 50, minus)
	fmt.Println(v)

	fun := choose(2)(18, 6)

	fmt.Println(fun)

}
