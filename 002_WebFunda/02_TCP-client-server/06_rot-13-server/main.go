package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	svr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer svr.Close()
	log.Println("Started ROT-13 server on port 8080")
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
	err := conn.SetDeadline(time.Now().Add(50 * time.Second))
	if err != nil {
		log.Println("Unable to set Connection Timeout")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)
		s := fmt.Sprintf("%s - %s\n", ln, r)
		conn.Write([]byte(s))
		log.Println(s)
	}
	log.Println("Connection Terminated....")
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
