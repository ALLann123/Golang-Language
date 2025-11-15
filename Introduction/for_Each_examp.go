package main

import "fmt"

func main(){
	var tools []string
	var t string
	for i:=0;i<3;i++{
		fmt.Println("Enter tool Name>> ")
		fmt.Scan(&t)
		tools=append(tools, t)
	}

	//display
	for _, tool:=range tools{
		fmt.Printf("Tool: %v \n", tool)
	}
}

/*
Enter tool Name>> 
burp
Enter tool Name>> 
msfconsole
Enter tool Name>> 
nmap
Tool: burp       
Tool: msfconsole 
Tool: nmap 
*/