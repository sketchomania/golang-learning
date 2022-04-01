package main

import (
	"encoding/json"
	"fmt"
)

// course is of private type we don't want to export it
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc321", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 499, "LearnCodeOnline.in", "bcda321", []string{"full-stack", "js", "mern"}},
		{"VeuJs Bootcamp", 249, "LearnCodeOnline.in", "let321go", []string{"web-dev", "community", "js"}},
		{"TypeScript Bootcamp", 199, "LearnCodeOnline.in", "let8750", nil},
	}

	//package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	// finalJson, err := json.MarshalIndent(lcoCourses, "sketch", "\t")
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Println("%s\n", finalJson)
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	JsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",
		"Price": 499,
		"website": "LearnCodeOnline.in",
		"tags": ["full-stack", "js", "mern"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(JsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(JsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(JsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		// order is not gurented in key value | map | # key
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
