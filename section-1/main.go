package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. String
	var name string = "muzaffar"
	fmt.Println(name)
	fmt.Printf("variable name is of type: %T\n", name)

	// 2. Bool
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable isLoggedIn is of type: %T\n", isLoggedIn)

	// 3. Int
	var number int = 30
	fmt.Println(number)
	fmt.Printf("variable number is of type: %T\n", number)

	// 4. Float
	var floatingNumber float64 = 30.99999999
	fmt.Println(floatingNumber)
	fmt.Printf("variable floatingNumber is of type: %T\n", floatingNumber)

	// Implicit types that lexer give us.
	var website = "freecodecamp.org"
	var (
		num1 = 2
		num2 = 4
	)
	fmt.Println(website)
	fmt.Println(num1 + num2)

	// no var style
	noOfUsers := 300
	fmt.Println(noOfUsers)

	// constant's
	const Pi = 3.14 // notice Pi first letter is Capital that makes it public
	fmt.Println(Pi)
	fmt.Printf("constant Pi is of type: %T\n", Pi)

	// Refer to comma-ok syntax used for error handling in go:
	fmt.Println("\nEnter a number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println(input)
	fmt.Printf("Type of input is: %T\n", input)

	// Type conversion
	convertedNum, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(convertedNum)
		fmt.Printf("Type of converted input: %T\n", convertedNum)
	}
}
