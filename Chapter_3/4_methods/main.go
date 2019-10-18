package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}

func main() {
	p := person{name: "Tom", age: 23}
	p.print()
	p.eat("бургер")
}
