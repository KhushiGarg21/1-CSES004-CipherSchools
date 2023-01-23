package main

import "fmt"

func main() {
	v := 3
	rec(v)
}
func rec(num int) {
	if num == 0 {
		return
	}
	if num%2 == 0 { //0 1 1
		rec(num - 1) // rec(1)
		fmt.Println(num - 1)
	} else {
		rec(num - 1) // rec(2) rec(0)
		fmt.Println(num - 1)
	}
	fmt.Println(num - 1)
}

// func fact(number int) int {
// 	if number == 1 && number == 0 {
// 		return 1
// 	}
// 	if number < 0 {
// 		return -1
// 	}
// 	result := number + fact(number-1)
// 	return result
// }
