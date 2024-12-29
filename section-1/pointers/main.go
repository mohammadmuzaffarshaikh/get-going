package main

import "fmt"

func main() {
	fmt.Println("Pointers:")
	num := 24
	var ptr *int = &num
	fmt.Println("Pointer address: ", ptr)
	fmt.Println("Pointer value: ", *ptr)

	*ptr = *ptr + 1
	fmt.Println("New value is: ", num)
}
