package main

import (
	"io"
	"log"
	"net"
)

//echo is a handler function that simply echoes recieved data.
//old echo function
/*
func echo(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store recieved data.
	b := make([]byte, 512)
	for {
		//Recieve data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected Error")
			break
		}
		log.Printf("Recieved %d bytes: %s\n", size, string(b))

		//Send data via conn.Write
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}
*/

//New Echo function (better)
func echo(conn net.Conn) {
	defer conn.Close()

	//Method 1
	/*
		reader := bufio.NewReader(conn)
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Unable to read data")
		}
		log.Printf("Read %d bytes: %s", len(s), s)

		log.Println("Writing data")
		writer := bufio.NewWriter(conn)
		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}
		writer.Flush()
	*/

	//Method 2
	// Copy data from io.Reader to io.Writer via io.Copy()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}

func main() {
	//Bind to TCP port 8080 on all interfaces.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:8080")
	for {
		//Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Recieved connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		//handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}
