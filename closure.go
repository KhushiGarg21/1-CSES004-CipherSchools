package main

import "fmt"

func main() {
	var m func(int) int //here function is stored as a value to a variable
	m = func(x int) int {
		fmt.Println("in a function 1", x)
		return 90
	}
	g(m, 8)

	j := func(x int) int {
		fmt.Println("in a function 2", x)
		return 20
	}(100)
	fmt.Println(j) //20
}
func g(m func(int) int, u int) { //passing function as a parameter to another function ie called closure
	m(244)
}
