package main

import "fmt"

// PrintLength is func for printing length
func PrintLength(str string) {
	l := len(str)

	printLine(l)

	fmt.Println(str)
	fmt.Println("Length:", l)

	printLine(l)
}

func printLine(length int) {
	for i := 0; i < length; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
