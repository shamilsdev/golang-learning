package main

import "fmt"

func main() {
	var people = map[string]int{"Tom": 1, "Bob": 2, "Alan": 3, "Sam": 4}
	fmt.Println(people)

	fmt.Println(people["Sam"])

	if val, ok := people["Smith"]; ok {
		fmt.Println(ok, val)
	}
}
