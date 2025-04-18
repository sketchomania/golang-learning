package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjk34dfio"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutorials",
		RawPath: "user=sketch",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
