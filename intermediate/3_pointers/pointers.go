package main

import "fmt"

func main() {
	//var ptr *int
	// Note:- Pointers arithmetic similar to C is not supported by GO.

	var a int = 10
	var ptr *int

	fmt.Println("Value stored after declaration -> ", ptr)

	ptr = &a

	fmt.Println("Value of a ->", a)
	fmt.Println("Value of Address stored in pointer ->", ptr)
	fmt.Println("Value stored at address inside pointer ->", *ptr)

	*ptr = 25
	fmt.Println("Value of a ->", a)
	fmt.Println("Value of Address stored in pointer ->", ptr)
	fmt.Println("Value stored at address inside pointer ->", *ptr)
}
