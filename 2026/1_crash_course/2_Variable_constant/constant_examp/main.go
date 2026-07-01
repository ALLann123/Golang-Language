package main

import "fmt"

//this is our constant value that will not change
const Pi = 3.14159

func main() {
	radius := 21

	//calculate
	fmt.Printf("Calculate the Area of the Circle:\n")

	//convert the radius to a floating point
	area := Pi * float64(radius) * float64(radius)

	fmt.Println("Circle Area is: ", area)
}
