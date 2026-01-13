package main

import (
	"fmt"
	"net"
	"sync"
)

// worker function processes ports from a channel and scans them
// ports: buffered channel containing port numbers to scan
// wg: pointer to WaitGroup to signal when scanning is complete
func worker(ports chan int, wg *sync.WaitGroup) {
	// Range over the channel - continuously receives port numbers until channel closes
	for p := range ports {
		// Construct target address for scanning
		address := fmt.Sprintf("192.168.1.118:%d", p)

		// Attempt TCP connection to check if port is open
		conn, err := net.Dial("tcp", address)

		if err == nil {
			// Port is open - close connection and report
			conn.Close()
			fmt.Printf("%d open\n", p)
		}
		// If port is closed, do nothing (don't print anything)

		// Signal that this port has been scanned
		wg.Done()
	}
	// Worker exits when channel closes
}

func main() {
	// Create buffered channel with capacity 100
	// This acts as a job queue for the worker pool
	ports := make(chan int, 100)

	// Create WaitGroup to track completion of all port scans
	var wg sync.WaitGroup

	// Create worker pool - start 100 concurrent scanners
	// Each worker runs in its own goroutine
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	// Feed port numbers 1-1024 into the channel for workers to process
	for i := 1; i <= 1024; i++ {
		// Increment WaitGroup counter for each port to scan
		wg.Add(1)

		// Send port number to workers via channel
		// Will block if all 100 workers are busy and channel is full
		ports <- i
	}

	// Wait for all port scans to complete
	// Blocks until all 1024 wg.Done() calls have been made
	wg.Wait()

	// Close the channel - signals workers to exit their range loops
	close(ports)
}

/*
	cmd>> go run main.go
80 open
53 open
23 open
25 open
22 open
21 open
111 open
139 open
445 open
513 open
514 open
512 open

*/
