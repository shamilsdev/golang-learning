package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	bob := person{name: "Bob", age: 23}
	fmt.Println(bob)

	var bobPointer *person
	bobPointer = &bob
	bobPointer.age = 33
	fmt.Println(*bobPointer)

	var agePointer = &bob.age
	*agePointer = 45
	fmt.Println(bob)
}
