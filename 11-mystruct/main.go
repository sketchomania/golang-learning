package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang;
	// NO super, NO child OR parent in golang

	sketchomania := User{"Vaibhav", "vaibhavvast6529@gmail.com", true, 12}
	fmt.Println(sketchomania)
	fmt.Println(sketchomania.Email)
	fmt.Printf("sketchomania details are; %v\n", sketchomania)
	fmt.Printf("sketchomania details are; %+v\n", sketchomania)
}

type User struct {
	// "Name" must start from capital letter if you want it to be excessible outside
	Name   string
	Email  string
	Status bool
	Age    int
}
