//Basic Reader and Writer in Go

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// myReader defines an io.Reader to read from stdin
type myReader struct{}

//Read reads data from stdin
func (myotherReader *myReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// myWriter defines an io.Writer to write to Stdout
type myWriter struct{}

// Write writes data to Stdout
func (myOtherWriter *myWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	//Instansiate Reader and Writer
	var (
		reader myReader
		writer myWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")

	}
	//Alternative to code below

	/*
		//create a buffer to hold input /output
		input := make([]byte, 4096)

		//Use reader to read input
		s, err := reader.Read(input)
		if err != nil {
			log.Fatalln("Unable to read data")
		}
		fmt.Printf("Read %d bytes from stdin\n", s)

		//Use Writer to write output
		s, err = writer.Write(input)
		if err != nil {
			log.Fatalln("Unable to write data")
		}
		fmt.Printf("Wrote %d bytes to stdout\n", s)
	*/

}
