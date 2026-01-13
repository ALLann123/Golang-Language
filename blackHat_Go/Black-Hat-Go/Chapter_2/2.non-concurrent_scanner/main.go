package main

import (
	"fmt"
	"net"
)

func main() {
	//Make a for loop to scan all ports from 1 to 1024
	for i := 1; 1 <= 1024; i++ {
		//convert the value "i" which is an integer to a string using sprintf as the net.dial() requires only strings
		address := fmt.Sprintf("scanme.nmap.org:%d", i)

		//check if there is a TCP connection to see if opened
		conn, err := net.Dial("tcp", address)

		//if nothing is returned the port is opened if not it is closed/filtered so we use 'continue' to start the loop from here and skip everything
		if err != nil {
			//Port is closed or filtered
			continue
		}
		//FINishing our connections
		conn.Close()

		fmt.Printf("%d open \n", i)
	}
}

/*
	cmd>> go run main.go
22 open
80 open
*/
