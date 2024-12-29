package main

import "fmt"

func main() {
	fmt.Println("Maps in go:")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["py"] = "Python"
	languages["ru"] = "Rust"

	fmt.Println(languages)

	// Delete
	delete(languages, "py")
	fmt.Println(languages)

	for _, value := range languages {
		fmt.Printf("%v\n", value)
	}
}
