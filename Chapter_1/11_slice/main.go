package main

import "fmt"

func main() {
	var users1 = [3]string{"Tom", "Alice", "Kate"}
	fmt.Println(users1)

	users2 := []string{"Tom", "Alice", "Kate"}
	fmt.Println(users2)

	var intSlice []int = make([]int, 3)
	fmt.Println(intSlice)
	intSlice[0] = 12
	intSlice[1] = 122
	intSlice[2] = 123
	fmt.Println(intSlice)

	var users []string
	users = append(users, "Alice")
	users = append(users, "Mark")
	users = append(users, "Newton")
	fmt.Println(users)
	fmt.Println("-----------------------")

	var users3 []string = make([]string, 3)
	users3[0] = "first"
	users3[1] = "second"
	users3[2] = "third"
	fmt.Println(users3)

	v := users3[0:2]
	v1 := users3[1:3]
	fmt.Println(v)
	fmt.Println(v1)

	fmt.Println("-----------------------")
	fmt.Println(users3)
	users3 = append(users3[:1], users3[2:]...)
	fmt.Println(users3)
}
