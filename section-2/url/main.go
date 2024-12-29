package main

import (
	"fmt"
	"net/url"
)

var urlOne string = `https://mycode.in/learn/go?page=3&id=222`

func main() {
	result, _ := url.Parse(urlOne)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	for _, val := range qparams {
		fmt.Println(val)
	}

	// create your own url

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "me.com",
		Path:     "/mme",
		RawQuery: "user=123",
	}

	fmt.Println(partsOfUrl.String())
}
