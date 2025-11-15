package main

import "fmt"

func main(){
	//Prints what the user enters
	var userInput string
	for{
		fmt.Println("User: ")
		fmt.Scan(&userInput)

		if userInput=="quit"{
			break
		}else{
			fmt.Printf("AI(Reflect): %v", userInput)
			fmt.Println()
		}

	}
}