package main

import "fmt"

func main() {
	defer fmt.Println("work")
	defer fmt.Println("yo")
	fmt.Println("hello")
	myDefer()
}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

/*
	what defer do is basically put the block of code in stack, where first in is last out
	so above output will be:
	hello
	5
	4
	3
	2
	1
	yo
	work
*/
