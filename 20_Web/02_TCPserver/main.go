package main

import (
	"net"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func handle(conn net.Conn) {
	// Buffered Scanner
	scanner := bufio.NewScanner(conn)
	// Make sure at the end we close the connection
	defer conn.Close()
	//Wait till scanning goes
	for scanner.Scan() {
		// Get the Text that was sent
		line := scanner.Text()
		// Just log it
		log.Println(line)
		// Terminate chat in case exit is received from other end
		if strings.ToLower(line) == "exit" {
			break // This helps to keep the Host / Server running - No need to Termination
		}
		// Add the Prepended text
		line = fmt.Sprintln(" From Server: ", line)
		// Send it out the Connection
		fmt.Fprint(conn, line)
	}
}

func main() {
	link, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server Started on port 9000")
	defer link.Close()

	for {
		conn, err:=link.Accept()
		if err != nil {
			panic(err)
		}
		handle(conn)
		/*
		// Simpler Way to make a echo server, but each time connection would close automatically
		io.Copy(conn, conn)
		conn.Close()*/

	}
}
