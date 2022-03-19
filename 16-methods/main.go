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
	sketchomania.GetStatus()
	sketchomania.NewMail()
	fmt.Printf("sketchomania details are; %+v\n", sketchomania)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
