package main

import "fmt"

func main(){
	/*
	To Understand if....else statements we get port number and get its Service
	21-FTP, 22-SSH, 23-Telnet, 25-SMTP, 80-HTTP, 443-HTTPS, 512-ModBus
	*/

	var userInput string
	fmt.Println("*********************")
	fmt.Println("     PORT ID CLI         ")
	fmt.Println("*********************")
	for{
		fmt.Println("Enter Port>> ")
		fmt.Scan(&userInput)
		if userInput == "quit"{
			fmt.Println("GoodBye!!")
			break
		}else if userInput=="21"{
			fmt.Printf("Port is FTP\n")
		}else if userInput=="22"{
			fmt.Printf("Port is SSH\n")
		}else if userInput=="23"{
			fmt.Printf("Port is Telnet\n")
		}else if userInput=="25"{
			fmt.Printf("Port is SMTP\n")
		}else if userInput=="80"{
			fmt.Printf("Port is HTTP\n")
		}else if userInput=="443"{
			fmt.Printf("Port is HTTPS\n")
		}else if userInput=="512"{
			fmt.Printf("Port is ModBus\n")
		}else{
			fmt.Println("UNKNOWN!!")
		}
	}
}

/*
*********************
     PORT ID CLI
*********************    
Enter Port>>
21
Port is FTP  
Enter Port>> 
22
Port is SSH  
Enter Port>> 
23
Port is Telnet
Enter Port>>  
25
Port is SMTP 
Enter Port>> 
*/
