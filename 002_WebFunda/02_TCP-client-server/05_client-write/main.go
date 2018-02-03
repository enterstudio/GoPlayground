package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	io.WriteString(conn, "Client Writing to Server at "+time.Now().Format(time.RFC3339Nano))
}
