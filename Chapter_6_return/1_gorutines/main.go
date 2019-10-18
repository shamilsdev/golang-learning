package main

import "fmt"

func calc(x, t int) int {
	return x + t
}

func main() {
	for index := 0; index < 10; index++ {
		go fmt.Println(calc(32, index*23))
	}
}
