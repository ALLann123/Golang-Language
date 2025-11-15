package main

import "fmt"

func main(){
	//declare variable
	var IP string
	fmt.Println("Enter Target>> ")
	fmt.Scan(&IP)

	fmt.Printf("[+]Ransomware target %v is executing code...", IP)

}