package main

import "fmt"

func main() {
	fmt.Println("-------------Arrays----------------------------")
	var numbers [10]int
	numbers[0] = 0
	numbers[1] = 1
	numbers[2] = 2

	var numbers2 [2]string = [2]string{"first", "second"}
	fmt.Println(numbers2)

	var numbers3 = [2]string{"third", "fourth"}
	fmt.Println(numbers3)

	numbers4 := [2]string{"fifth", "sixth"}
	fmt.Println(numbers4)

	fmt.Println("-------------Indexes---------------------------")
	fmt.Println("numbers[2]", numbers[2])
	fmt.Println("numbers[1]", numbers[1])
}
