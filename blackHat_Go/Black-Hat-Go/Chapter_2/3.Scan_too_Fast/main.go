package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	//Create a wait group to track completion of all goroutines
	var wg sync.WaitGroup

	//scan ports 1 through 1024
	for i := 1; i <= 1024; i++ {
		//Increament waitGroup conuter for each goroutines we're about to launch
		wg.Add(1)

		//launch a goroutine to check each port
		//we pass 'i' as parameter to avoid closure variable capture issues
		go func(j int) {
			//Decrement WaitGroup when goroutines completes
			defer wg.Done()

			//construct the target address with port number--> Convert integer port to string
			address := fmt.Sprintf("192.168.1.118:%d", j)

			//Attempt TCP connection to the address
			//net.Dial returns a connection or error
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return //Exit go routines silently for closed ports
			}

			//If connection succedded, close it(we wanred to check if open)
			conn.Close()

			//print the open port number
			fmt.Printf("%d open\n", j)
		}(i) //Pass current port number i as input j
	}

	//Wait for all go routines to complete before exiting main()
	wg.Wait()
}

/*
Target scanme.org:
	cmd>. go run main.go
22 open
80 open

Target Metasploitbale:
	cmd>> go run main.go
21 open
22 open
111 open
25 open
23 open
53 open
139 open
512 open
513 open
514 open
445 open
80 open
*/
