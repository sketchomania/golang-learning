package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	// map[key]value
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Guby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS stnads for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang
	// %v is for value

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
		//fmt.Println("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
