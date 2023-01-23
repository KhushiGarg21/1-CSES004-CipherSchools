package main

import "fmt"

func main() {
	result, x := a()
	fmt.Println(result, x)

	result2 := b(1, 2, 3, 4, 5, 6)
	fmt.Println(result, x, result2)

	k := new(int)
	*k = 21
	var d = 10
	pointer(&d)
	fmt.Println(d) // pointer will work, and the updated value of d is printed

}
func a() (int, string) {
	return 10, ""
}
func b(args ...int) bool {
	for _, v := range args {
		fmt.Println(v)
	}
	return true
}
func pointer(k *int) { // here k is pointer and we are declaring its type
	*k = 202
}
