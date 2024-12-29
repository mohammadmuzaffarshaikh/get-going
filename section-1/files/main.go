package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files:")
	content := "my file is reading and already writen this in it."

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	} else {
		length, _ := io.WriteString(file, content)
		fmt.Printf("Length of file is %v\n", length)
		defer file.Close()
	}
	readFile(file.Name())
}

func readFile(file string) {
	databyte, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
