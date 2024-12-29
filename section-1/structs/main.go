package main

import "fmt"

func main() {
	muzaffar := User{"Muzaffar", "muzaffar@gmail.com", 22}

	fmt.Println(muzaffar)

	fmt.Printf("Detail of user: %+v\n", muzaffar)

	fmt.Printf("Name is %v\n", muzaffar.Name)
}

type User struct {
	Name  string
	Email string
	Age   int
}
