package main

import "fmt"

func main() {
	i := 0

	for {
		if i > 5 {
			break //exit the loop
		}

		fmt.Println("Infinite loop iteration: ", i)
		i++
	}
}
