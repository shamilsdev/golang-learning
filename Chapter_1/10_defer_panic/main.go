package main

import "fmt"

func main() {

	defer fmt.Println("1 - gooolang!")
	defer fmt.Println("2 - gooolang!")
	defer fmt.Println("3 - gooolang!")
	fmt.Println("Hello")
	panic("Custom error")
}

func finish() {
	panic("Custom error")
}
