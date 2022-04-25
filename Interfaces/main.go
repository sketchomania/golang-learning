package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type Rect struct {
	lenght, width float64
}

func (r Rect) area() float64 {
	return r.lenght * r.width
}
func (r Rect) perim() float64 {
	return 2*r.lenght + 2*r.width
}

// Spliting shape interface into two interfaces shape1 and shape2
type Shape1 interface {
	// shape1 require only perim() method
	perim() float64
}

type Shape2 interface {
	// shape2 require only area() method
	area() float64
}

type Shape interface {
	// Circle is a shape and Rectangle is a shape
	// area() float64
	// perim() float64
	// shape require both area() and perim() method
	// shape is compose of interface shape1 and shape2
	Shape1
	Shape2
}

// func measure(s Shape) {
// 	fmt.Println(s)
// 	fmt.Println("Area: ", s.area())
// 	fmt.Println("Parameter: ", s.perim())
// }

// func measureCircle(s Shape) {
// 	fmt.Println(s)
// 	fmt.Println("Area: ", s.area())
// 	fmt.Println("Parameter: ", s.perim())
// }

// func measureRect(s Shape) {
// 	fmt.Println(s)
// 	fmt.Println("Area: ", s.area())
// 	fmt.Println("Parameter: ", s.perim())
// }

func main() {
	fmt.Println("Welcome to learning interfaces in golang")
	// intro()

	// c := Circle{radius: 1}
	// r := Rect{lenght: 5, width: 2}
	// // fmt.Printf("[c] Value = %+v, Type = %T\n", c, c)
	// fmt.Printf("[r] Value = %+v, Type = %T\n", r, r)

	// measureCircle(c)
	// measureRect(r)

	// measure(c)
	// measure(r)

	var s Shape = Circle{radius: 1}
	fmt.Printf("[s] Value = %+v, Type = %T\n", s, s)
	// fmt.Printf("Value = %+v, Type = %T\n", s.radius, s.radius) // err: shape does not have field or method radius

	// converting concrete type from interfaces
	// converting shape s into circle
	// c := s.(Circle)
	// fmt.Printf("Value = %+v, Type = %T\n", c, c)
	// fmt.Printf("Value = %+v, Type = %T\n", c.radius, c.radius)

	// c := s.(Rect)
	// fmt.Printf("Value = %+v, Type = %T\n", c, c) // panic: main.Shape is main.Circle, not main.Rect

	// run 1(no err) or 2(err)
	// c, ok := s.(Circle)  //1
	// // c, ok := s.(Rect)  //2
	// if ok {
	// 	fmt.Printf("Value = %+v, Type = %T\n", c, c)
	// } else {
	// 	fmt.Println("Cannot convert")
	// }
}

func intro() {
	fmt.Println("An interface type is defined as a set of methods signatures.")
	fmt.Println("A value of interface type can hold any value that implememnts those methods.")
	fmt.Println("A type implements an interface by implementing its methods.")
	fmt.Println("There is no explicit declaration of intent, no implements keyword.")
	fmt.Println("Interfaces are composite interfaces, In go we can compose interfaces of other interfaces")
}
