package main

import "fmt"

func main(){
	//get 3 IPs
	var IP []string
	var j string
	
	for i:=0; i<3;i++{
		fmt.Println("Enter IP>> ")
		//Pause program to get user input
		fmt.Scan(&j)
		//Add our values to the SLICE
		IP=append(IP, j)
	}

	for y:=0; y<3;y++{
		fmt.Printf("[+]Running ransomware on %v.... \n", IP[y])
	}

	fmt.Println()
	fmt.Println("Sleeping!!")

}