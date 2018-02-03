package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	log.Println("TCP Server")
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	log.Println("TCP Server Started on port 8080")
	defer listner.Close()
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprintln(conn, "How are you doing today ?")
		fmt.Fprintf(conn, "Meaning of Life %d\n", 108)
		conn.Close()
	}
}
