package main

import (
	"fmt"

	"./computation"
)

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))

	fmt.Println(computation.Calc(5, 5))
}
