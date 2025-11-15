package main

import "fmt"

func main(){
	//declare variables
	var length float64
	var width float64
	var area float64
	var perimeter float64

	//get user input
	fmt.Println("Enter length: ")
	fmt.Scan(&length)

	fmt.Println("Enter Width: ")
	fmt.Scan(&width)

	//calculate area
	area=length*width

	//calculate perimeter
	perimeter=(length+width)*2

	fmt.Printf("The Area is %v and the Perimeter is %v \n", area, perimeter)
}