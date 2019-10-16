package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibbonachi(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}

func main() {
	v := factorial(5)
	fmt.Println(v)

	f := fibbonachi(15)
	fmt.Println(f)
}
