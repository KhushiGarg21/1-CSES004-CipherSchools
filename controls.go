package main

import "fmt"

func main() {
	fmt.Println("ENter a number")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 10:
		input = 400
		fmt.Println(input)
	case 11:
		input = 100
		fmt.Println(input)
	case 12:
		input = 100
		fmt.Println(input)

	default:
		input = 200
		fmt.Println(input)
	}
}
