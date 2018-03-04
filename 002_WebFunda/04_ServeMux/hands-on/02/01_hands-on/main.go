package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	svr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Started TCP Server at 8080")
	for {
		conn, err := svr.Accept()
		if err != nil {
			log.Println(err)
		} else {
			go handle(conn)
		}
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	err := conn.SetDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	io.WriteString(conn, "I see you connected.")
}
