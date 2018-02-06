package main

import (
	"bufio"
	"bytes"
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
	log.Println("Server Started on port 8080")

	for {
		conn, err := svr.Accept()
		if err != nil {
			log.Println(err)
		} else {
			log.Println(" New Connection -------------")
			log.Println(conn.RemoteAddr().String())
			go handleHTTP(conn)
		}
	}
}

func handleHTTP(conn net.Conn) {
	defer conn.Close()
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Could not set Deadline", err)
	}

	// Handle the Request
	requestHTTP(conn)

	// Handle the Response
	responseHTTP(conn)

	log.Println("--------  Connection Terminated.....")
}

// requestHTTP does the handling of the Incoming request on the connection
func requestHTTP(conn net.Conn) {
	var uri string
	var method string
	var contentType string
	headers := make(map[string]string, 20)
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			break
		}
		log.Println("  ~ ", ln)

		// Break the Line into fields
		fl := strings.Fields(ln)
		// Process the Request
		if i == 0 && len(fl) > 2 {
			method = fl[0]
			uri = fl[1]
		}
		if strings.Contains(ln, "Accept:") && len(fl) > 1 {
			contentType = fl[1]
		}
		if i != 0 && len(fl) > 1 {
			key := strings.Trim(fl[0], ":-")
			headers[key] = strings.TrimSpace(strings.Join(fl[1:], ""))
		}

		// Increment the Counter
		i++
	}
	log.Println("> Method\t:", method)
	log.Println("> URI\t:", uri)
	log.Println("> Content-Type:", contentType)
	log.Println("> Headers:", headers)
}

func responseHTTP(conn net.Conn) {
	body := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>Hello Web Page</title>
</head>
<body>
	<h1>Hare Krishna</h1>
</body>
</html>
`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
}

// Modularizing above code into a Context based execution

// ProcCtx is a data structure to store the TCP
//  Connection information for an HTTP request
type ProcCtx struct {
	Connection net.Conn                  // Handle of the TCP Connection
	Head       bytes.Buffer              // Request Storage
	URI        string                    // URI in the Current Request
	Method     string                    // Method of the Current Reqest
	Headers    map[string]string         // Headers processed from the Current Request
	Lines      int                       // Number of Lines in the Current Request
	Handles    map[string]func(*ProcCtx) // List of Handlers for specific URI
	Debug      bool                      // In case Debug mode is Enabled to do Logging
}

// NewProcCtx is function that creates a new ProcCtx value using the
// incoming connection handle.
func NewProcCtx(conn net.Conn) *ProcCtx {
	return &ProcCtx{
		Connection: conn,
	}
}

func (p *ProcCtx) processRequest() {
	p.Headers = make(map[string]string, 20)
	p.Lines = 0 // Which Line of Processing we are On
	if p.Debug {
		log.Println("---- DEBUG log for processRequest() [Begin] ----")
	}
	// Process the Incoming Stream
	scanner := bufio.NewScanner(p.Connection)
	for scanner.Scan() && p.Lines < 20 {
		ln := scanner.Text()
		p.Head.WriteString(ln)
		if len(ln) == 0 {
			break
		}

		if p.Debug {
			log.Printf(" Line -%d-: %s\n", p.Lines, ln)
		}

		// Break the Line into fields
		fl := strings.Fields(ln)

		// Process the Request
		if p.Lines == 0 && len(fl) > 2 {
			p.Method = fl[0]
			p.URI = fl[1]
		}
		if p.Lines != 0 && len(fl) > 1 {
			key := strings.Trim(fl[0], ":-")
			p.Headers[key] = strings.TrimSpace(strings.Join(fl[1:], ""))
			if p.Debug {
				log.Printf(" Fields: \n  %v : %v\n", key, p.Headers[key])
			}
		}

		// Increment the  Line Counter
		p.Lines++
	}

	if p.Debug {
		log.Printf(" URI: %s", p.URI)
		log.Printf(" Method: %s", p.Method)
		log.Printf(" Total Lines: %d", p.Lines)
		log.Println("---- DEBUG log for processRequest() [ End ] ----")
	}
}
