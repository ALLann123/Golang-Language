package main

import (
	"fmt"
	"log"
	"os"
)

// FooReader defines an io.Reader to read from stdin
type FooReader struct{}

// Read reads data from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// Foowriter defines an io.Writer to write to stdout
type Foowriter struct{}

// Write writes data to Stdout.
func (fooWriter *Foowriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	//Instantiate reader and writer
	var (
		reader FooReader
		writer Foowriter
	)

	//create buffer to hold i/o
	input := make([]byte, 4096)

	//use reader to read input
	s, err := reader.Read(input)

	if err != nil {
		log.Fatalln("Unable to read data")
	}

	fmt.Printf("Read %d bytes from stdin\n", s)

	//use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}

/*
	cmd>> go run main.go
in > hello world
Read 13 bytes from stdin
out> hello world
*/
