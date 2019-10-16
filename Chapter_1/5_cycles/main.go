package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	var l = 0
	for ; l < 9; l++ {
		fmt.Print(l, " ")
	}
	fmt.Println()

	var n = 0
	for n < 8 {
		fmt.Print(n, " ")
		n++
	}
	fmt.Println()

	var p = 0
	for {
		fmt.Print(p, " ")
		p++
		if p > 50 {
			break
		} else if p > 20 {
			continue
		} else {
			fmt.Print("|")
		}
	}
	fmt.Println()

	for x := 1; x < 13; x++ {
		for y := 1; y < 13; y++ {
			fmt.Print(x*y, "\t\t")
		}
		fmt.Println()
	}

	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, n := range numbers {
		fmt.Println(i+1, ")", "\t", n)
	}

}
