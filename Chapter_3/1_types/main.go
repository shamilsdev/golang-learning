package main

import "fmt"

type kilometer int
type mile int

func distanseTo(m mile) {
	fmt.Println(m, "MILE")
}

type library []string

func printLib(lib []string) {
	for i, val := range lib {
		fmt.Println(i, val)
	}
}

func main() {
	var distance mile = 5
	distanseTo(distance)

	fmt.Println("----------------------")

	var lib library = library{"Voina", "Mir"}
	printLib(lib)
}
