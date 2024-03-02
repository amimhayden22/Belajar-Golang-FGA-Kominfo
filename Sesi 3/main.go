package main

import "fmt"

func main() {
	var word string = "hello"

	_ = word

	var romanian string = "Mânca"

	// fmt.Println(len(romanian))

	for i := 0; i < len(romanian); i++ {
		fmt.Println(romanian[i], string(romanian[i]))
	}

	// for i := 0; i < count; i++ {
		
	// }

	// char := word[0]

	// var char string = string(word[0:3])

	// fmt.Println(char)

	// for i := 0; i < len(word); i++ {
	// 	fmt.Println(string(word[i]))
	// }

	// _ = char

	// fmt.Println(word[0])
}
