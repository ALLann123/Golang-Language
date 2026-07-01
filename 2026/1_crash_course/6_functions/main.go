package main

import "fmt"

func main() {
	//main executiion of a program is here
	num_1 := 23
	num_2 := 12

	// call the addition function
	result := add(num_1, num_2)

	//print output
	fmt.Println("Total is: ", result)
}

//write our addition function and returns the sum. Accepts two integers
func add(a int, b int) int {
	return a + b
}
