package main

import (
	"fmt"
	"io"
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
	var (
		reader FooReader
		writer Foowriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

/*
	cmd>> go run main.go
in > Hacked
out> Hacked
in > Hey
out> Hey
in > ping
out> ping
in > ping google.com
out> ping google.com
*/
