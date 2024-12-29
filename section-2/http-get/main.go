package main

import (
	"fmt"
	"io"
	"net/http"
)

var url string = "https://daryl-ng.medium.com/why-you-should-avoid-ioutil-readall-in-go-e6be4de180f8"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Response is of type: %T\n", response)

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databyte)

	fmt.Println(content)
}
