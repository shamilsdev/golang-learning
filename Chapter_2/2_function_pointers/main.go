package main

import "fmt"

func main() {
	x := 45

	fmt.Println(x)

	changeParam(&x)

	fmt.Println(x)
}

func changeParam(p *int) {
	*p = *p * *p
}
