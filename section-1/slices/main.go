package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "mango"}
	fmt.Printf("Type of fruits is %T\n", fruits)

	// adding elements to slice
	fruits = append(fruits, "banana", "peach")
	fmt.Println(fruits)

	// Picking specific elements in range
	fruits = append(fruits[1:2])
	fmt.Println("now fruits: ", fruits)

	// another way to make slice
	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 3263
	highScores[2] = 1623
	highScores[3] = 990

	highScores = append(highScores, 240, 557)

	// sorting the slice
	sort.Ints(highScores)
	fmt.Println(highScores)

	// deleting a specific index element
	var index int = 2
	var courses = []string{"React", "Node", "JS", "GO", "Python", "Rust"}
	fmt.Println("Before removing element: ", courses)
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After deleting the value: ", courses)
}
