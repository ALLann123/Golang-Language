package main

import "fmt"

func main() {
	var length float64
	var width float64
	fmt.Println("Enter the Length: ")
	fmt.Scan(&length)
	fmt.Println("Enter the width: ")
	fmt.Scan(&width)

	area := calc_area(length, width)
	fmt.Printf("Area is %v \n", area)

	perimeter := calc_perimeter(length, width)
	fmt.Printf("Perimeter is %v \n", perimeter)
}

func calc_area(length, width float64) float64 {
	area := length * width
	return area
}

func calc_perimeter(length, width float64) float64 {
	perimeter := (length + width) * 2
	return perimeter
}
