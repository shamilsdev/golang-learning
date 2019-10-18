package main

import "fmt"

func main() {
	x := 45

	fmt.Println(x)

	changeParam(&x)

	fmt.Println(x)

	changeParam2(&x)

	fmt.Println(x)

	fmt.Println("------------------------------")

	val := new(int)

	fmt.Println(*val)

	fmt.Println("------------------------------")

	createPointer(10)

}

func changeParam(p *int) {
	*p = *p * *p
}

func changeParam2(p *int) {
	*p = *p * *p
}

func createPointer(x int) {
	fmt.Println(x)

	val := new(int)
	fmt.Println(val)

	*val = x
	fmt.Println(*val)
	fmt.Println(&val)
}
