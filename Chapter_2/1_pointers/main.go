package main

import "fmt"

func main() {
	var p int = 155

	u := &p

	*u = 200

	fmt.Println(p)

	fmt.Println(*u)

	t := &p

	if t != nil {
		fmt.Println(*t)
	}

	fmt.Println("--------------------------------------")

	x := new(string)

	fmt.Println(*x)

	*x = "string"

	fmt.Println(*x)
}
