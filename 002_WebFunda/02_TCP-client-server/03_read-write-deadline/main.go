package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	svr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer svr.Close()

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
	// Set Timeout for the
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection Timeout setting Error")
	}

	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		log.Println(ln)
		s := fmt.Sprintf("I heard you say: %s\n", ln)
		log.Println(s)
		conn.Write([]byte(s))
	}

	// Now Code Reaches here
	log.Println("Code got here. and connection Terminated...")
}
