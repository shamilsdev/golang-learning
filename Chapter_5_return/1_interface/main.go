package main

import "fmt"

type vehicle interface {
	move()
}

type Car struct{}

type Airplane struct{}

func (c Car) move() {
	fmt.Println("Car")
}

func (c Airplane) move() {
	fmt.Println("Airplane")
}

func main() {
	c := Car{}
	c.move()

	a := Airplane{}
	a.move()
}
