package main

import "fmt"

func main() {

	//1st method:-
	i := 10
	j := &i
	*j = 100
	fmt.Println(i, j)

	//2nd method:-

	var i1 int
	i1 = 10
	var j1 *int //declarations only, as there is nothing in j1
	fmt.Println("dec", j1)

	j1 = &i1
	*j1 = 100
	fmt.Println(i1, j1)

	//3rd method:-

	var i3 int
	i3 = 10
	j3 := new(int) //declaration + assigment( as there is some address stored at j3)
	fmt.Println("dec2", j3)

	j3 = &i3
	*j3 = 100
	fmt.Println(i3, j3)

	//4th method:-
	name := new(string)
	*name = "golang"
	fmt.Println(name)

}
