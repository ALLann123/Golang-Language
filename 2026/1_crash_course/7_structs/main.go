package main

import "fmt"

// define a Circle Struct---> similar to class in OOP
type Circle struct {
	//the value we accept
	radius float64
}

// Define a method to calculate the area of the circle
// Pass our struct as the input of our function
func (c Circle) area() float64 {
	return 3.14159 * c.radius * c.radius
}

func main() {
	//call our struct and passing in a floating value
	Areacircle := Circle{radius: 7.0}

	//print the result
	fmt.Println("Circle Area:", Areacircle)
}
