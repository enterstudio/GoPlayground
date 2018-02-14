package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type httpcontext struct {
	uri    string
	method string
}

func (h *httpcontext) request(c net.Conn) {
	scanIn := bufio.NewScanner(c)
	line := 0 // Indicates which line is in process
	for scanIn.Scan() {
		ln := scanIn.Text()

		// Exit upon an Empty line
		if len(ln) == 0 {
			break
		}

		log.Printf(" ~ %s", ln)

		fields := strings.Fields(ln)

		if line == 0 && len(fields) > 2 {
			h.method = fields[0]
			h.uri = fields[1]
		}

		line++
	}
}

func (h *httpcontext) response(c net.Conn) {
	switch h.uri {
	case "/":
		pageHome := `
		<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
		<html><head><title>Home Page</title></head>
		<body>
		<h1>Welcome</h1>
		<p>This is the Home Page</p>
		</body></html>
		`
		s := fmt.Sprintf("HTTP/1.1 200 Ok\r\n")
		s += fmt.Sprintf("Date: %v\r\n", time.Now().String())
		s += fmt.Sprintf("Server: %s\r\n", "Golang Server")
		s += fmt.Sprintf("Content-Length: %s\r\n", len(pageHome))
		s += fmt.Sprintf("Connection: close\r\n")
		s += fmt.Sprintf("Content Type: text/html\r\n")
		s += fmt.Sprintf("\r\n")
		s += pageHome

		log.Println(" RESPONSE PAGE: \r\n", s)

		c.Write([]byte(s))
	default:
		page404 := `
		<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
		<html><head><title>404 Not Found</title></head>
		<body>
		<h1>Not Found</h1>
		<p>The requested URL %s was not found on this server.</p>
		</body></html>
		`
		page404 = fmt.Sprintf(page404, h.uri)
		s := fmt.Sprintf("HTTP/1.1 404 Not Found\r\n")
		s += fmt.Sprintf("Date: %v\r\n", time.Now().String())
		s += fmt.Sprintf("Server: %s\r\n", "Golang Server")
		s += fmt.Sprintf("Content-Length: %s\r\n", len(page404))
		s += fmt.Sprintf("Connection: close\r\n")
		s += fmt.Sprintf("Content Type: text/html\r\n")
		s += fmt.Sprintf("\r\n")
		s += page404

		log.Println(" RESPONSE PAGE: \r\n", s)

		c.Write([]byte(s))

	} // End of Switch
}

func main() {
	srv, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Starting TCP server at port 8080")

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Println(err)
		} else {
			go func(conn net.Conn) {
				err := conn.SetDeadline(time.Now().Add(5 * time.Second))
				if err != nil {
					log.Panicln(err)
				}
				defer conn.Close()
				var h httpcontext

				h.request(conn)

				h.response(conn)
			}(conn)
		}
	}
}
