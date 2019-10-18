package main

import "fmt"

type contactInfo struct {
	phone string
	email string
}

type person struct {
	name    string
	age     int
	contact contactInfo
}

type node struct {
	val  int
	next *node
}

func main() {
	p := person{name: "Tom", age: 24, contact: contactInfo{phone: "beeline", email: "mail.ru"}}
	fmt.Println(p)

	fmt.Println("----------------------------")

	first := node{val: 1}
	second := node{val: 2}
	third := node{val: 3}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

}
