package main

import "fmt"

func main(){
	/*
	1. Get a list of 7 IP targets
	---> Store in an Array
	2. Run Ransomware on the targets
	*/
	var targetIPS [7]string

	//get user input
	for i:=0; i<7;i++{
		fmt.Println("Enter IP: ")
		fmt.Scan(&targetIPS[i])
	}

	//start attack
	for j:=0; j<7; j++{
		fmt.Printf("[+]Attacking %v with Ransomware......\n", targetIPS[j])
	}
}