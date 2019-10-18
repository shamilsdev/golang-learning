package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) run(s string) {
	(*p).age += 10
	fmt.Println((*p).name, (*p).age, s)
}

func (p person) runn(s string) {
	p.age += 10
	fmt.Println(p.name, p.age, s)
}

func do() {
	var bob person = person{"Bob", 26}
	fmt.Println(bob)
	var bobPoint = &bob
	bobPoint.run("10 км")
	fmt.Println(bob)
}

func doo() {
	var bob person = person{"Bob", 26}
	fmt.Println(bob)
	bob.runn("100 км")
	fmt.Println(bob)
}

func (p person) updateAge(newAge int) {
	p.age = newAge
}

func (p *person) updateAge2(newAge int) {
	(*p).age = newAge
}

func main() {
	doo()

	fmt.Println("----------------------------------")

	var tom = person{name: "Tom", age: 24}
	fmt.Println("before", tom.age)
	tom.updateAge(33)
	fmt.Println("after", tom.age)

	t := &tom
	(*t).age = 44
	tom.updateAge((*t).age)
	fmt.Println("after after", tom.age)

	t.updateAge2(66)
	fmt.Println("after after after", tom.age)
}
