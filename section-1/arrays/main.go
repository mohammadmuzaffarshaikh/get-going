package main

import "fmt"

func main() {
	var fruits [3]string // var fruits = [3]string{"apple","banana","mango"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Println(fruits)
	fmt.Println("length: ", len(fruits))
}
