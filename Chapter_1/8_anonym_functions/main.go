package main

import "fmt"

func action(x, y int, parameter func(int, int) int) {
	v := parameter(x, y)
	fmt.Println(v)
}

func selectF(i int) func(int, int) int {
	switch i {
	case 1:
		return func(x int, y int) int { return x + y }
	case 2:
		return func(x int, y int) int { return x - y }
	default:
		return func(x int, y int) int { return x * y }
	}
}

func square(x int) func() int {
	return func() int { x++; return x * x }
}

func main() {
	var f func(int, int) int = func(x, y int) int { return x + y }
	//	f := func(x, y int) int { return x + y }

	fmt.Println(f(1, 15))

	action(10, 43, f)

	fmt.Println(selectF(12)(123, 321))

	funk := square(5)

	fmt.Println(funk())

}
