package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Backend stuff:")
	const url = "http://localhost:3000"
	GettingData(url + "/get")
}

func GettingData(Url string) {
	response, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("content length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	responseString.Write(content)

	fmt.Println(responseString.String())

	// fmt.Println("Content: ", string(content)) // you can do this but the above code is how senior dev do it.
}
