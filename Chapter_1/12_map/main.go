package main

import "fmt"

func main() {
	var people = map[string]int{"Mike": 1, "Rob": 2, "Pike": 3}

	if val, ok := people["Mikes"]; ok {
		fmt.Println(ok, val)
	}

	var users = make(map[int]string)
	users[1] = "Sam"
	users[2] = "Smith"
	users[3] = "Jonas"
	fmt.Println(users)

	users[4] = "Nick"

	delete(users, 3)
	fmt.Println(users)

}
