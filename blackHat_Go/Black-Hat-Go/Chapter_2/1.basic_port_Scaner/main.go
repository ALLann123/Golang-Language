package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")

	//lets check for errors being none using "nil"
	if err == nil {
		fmt.Println("Connection Successful")
	}

}

/*
	cmd>> go run main.go
Connection Successful
*/
