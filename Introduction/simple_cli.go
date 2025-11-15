package main

import "fmt"

func main(){
	var userName string
	var password string
	//get username and password to login
	for{
		fmt.Println("Enter Username: ")
		fmt.Scan(&userName)

		if userName=="admin"{
			fmt.Println("Enter Password: ")
			fmt.Scan(&password)
			if password=="kali"{
				fmt.Println("Shell>> ")
				break
			}
		}else{
			fmt.Println("[-]WRONG USERNAME.....")
		}
	}

	
}