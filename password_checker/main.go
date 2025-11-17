package main

import "fmt"

func main() {
	var password string
	for {
		fmt.Println("Enter Password(quit to exit):")
		fmt.Scan(&password)
		if password == "quit"{
			break
		}

		//call function to check strength
		checkPassword(password)
	}

}

func checkPassword(password string) {
	if len(password) >= 8 {
		fmt.Printf("%v---> Password Strong⚓\n", password)
	} else {
		fmt.Printf("%v-- Weak Password!!\n", password)
	}
}

/*
- Get password from user
- Check strength i.e >8 chars is strong
--> else weak
*/
