package main

import "fmt"

func main() {
	a := 100
	b := 8
	if a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a = b")
	}

	switch a {
	case 1:
		fmt.Println(1)
		break
	case 2:
		fmt.Println(2)
		break
	case 3:
		fmt.Println(3)
		break
	case 8:
		fmt.Println(8)
		break
	case 10:
		fmt.Println(10)
		break
	default:
		fmt.Println("not found")
	}
}
