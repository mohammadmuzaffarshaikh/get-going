package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, Open Up!")
	case 2:
		fmt.Println("Dice value is 2, Move 2!")
	case 3:
		fmt.Println("Dice value is 3, Move 3!")
	default:
		fmt.Println("YO bro what's you doing?")
	}
}
