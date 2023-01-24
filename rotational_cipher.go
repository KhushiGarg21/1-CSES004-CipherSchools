// func main() {

// 	shift := 13
// 	var msg string = "TRRXF SBE TRRXF"
// 	create_dict()
// 	encrypt(shift, msg)
// 	decrypt(shift, msg)
// }
// func create_dict() {
// 	map1 := make(map[int]string)
// 	map2 := make(map[int]string)
// 	for key, value := range map1 {
// 		map1[char[64+i]] = i
// 	}
// 	for key, value := range map2 {
// 		map1[char[64+i]] = i
// 	}
// }
// func encrypt(shift int, msg string) {
// 	var cipher string = ""
// 	for i := 0; i < 27; i++ {
// 		map1[char[64+i]] = i
// 	}
// 	for i := 0; i < 27; i++ {
// 		map1[char[64+i]] = i
// 	}
// }
// func decrypt(shift int, msg string) {
// 	var dicipher string = ""

// }

package main

import "fmt"

func main() {
	fmt.Println(Rotate("omg"))

}

var map1 = map[int]string{
	1:  "a",
	2:  "b",
	3:  "c",
	4:  "d",
	5:  "e",
	6:  "f",
	7:  "g",
	8:  "h",
	9:  "i",
	10: "j",
	11: "k",
	12: "l",
	13: "m",
	14: "n",
	15: "o",
	16: "p",
	17: "q",
	18: "r",
	19: "s",
	20: "t",
	21: "u",
	22: "v",
	23: "w",
	24: "x",
	25: "y",
	26: "z",
}

func getKey(letter string) (key int) {
	for i, value := range map1 {
		if value == letter {
			key = i
			break
		}
	}
	return key
}
func Rotate(msg string) (encrypt string) {
	for _, value := range msg {
		shiftedKey := getKey(string(value)) + 5
		shiftedKey = shiftedKey % 26

		encrypt += map1[shiftedKey]
	}
	return encrypt
}
