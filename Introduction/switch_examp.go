package main

import "fmt"

func main(){
	/*
	Enter service to get attack information
	*/
	var service string
	for{
		fmt.Println("Open Service: ")
		fmt.Scan(&service)
		if service=="quit"{
			fmt.Println("Byee!!")
			break
		}
		switch service{
			case "FTP", "SSH":
				fmt.Printf("%v---> try Bruteforcing\n", service)
				break
			case "HTTP","HTTPS":
				fmt.Printf("%v -->burp suite\n", service)
				break
			case "SMTP":
				fmt.Printf("%v---> social engineering\n", service)
				break
			default:
				fmt.Println("Do your own Research!!")
		}
	}
}

/*
Open Service: 
FTP
FTP---> try Bruteforcing
Open Service:
SSH
SSH---> try Bruteforcing
Open Service:
SMTP
SMTP---> social engineering
Open Service:
HTTPS
HTTPS -->burp suite
Open Service:      
quit
Byee!!
*/