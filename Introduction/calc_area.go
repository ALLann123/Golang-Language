package main

import "fmt"

func main(){
	const pi float32=3.142
	var radius float32=7
	//specify the data type
	var area float32

	area=pi*radius*radius
	fmt.Printf("The Area is: %v\n", area)

	//print the data types
	fmt.Printf("pi= %T, radius= %T, area=%T \n", pi, radius, area)
}