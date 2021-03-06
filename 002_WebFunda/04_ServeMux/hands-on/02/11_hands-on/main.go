package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	_ "time"
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
	/*err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		conn.Close()
		log.Fatalln(err)
	}*/
	go serve(conn)
	log.Println("Code got here.")
	//io.WriteString(conn, "I see you connected.")
}

func serve(c net.Conn) {
	defer c.Close()
	uri := ""
	method := ""
	line := 0
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			break
		}
		if line == 0 {
			fld := strings.Fields(ln)
			uri = fld[1]
			method = fld[0]
		}
		//io.WriteString(c, ln)
		log.Printf("Text Received: %s\n", ln)
		line++
	}
	log.Printf("URL: %s\n", uri)
	log.Printf("METHOD: %s\n", method)
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	io.WriteString(c, "\r\n")
}
