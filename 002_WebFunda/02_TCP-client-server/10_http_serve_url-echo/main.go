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
	uri := requestHTTP(conn)

	// Handle the Response
	responseHTTP(conn, uri)

	log.Println("--------  Connection Terminated.....")
}

// requestHTTP does the handling of the Incoming request on the connection
func requestHTTP(conn net.Conn) string {
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
	return uri
}

func responseHTTP(conn net.Conn, uri string) {
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
	<h2>URI : %s</h2>
</body>
</html>
`
	body = fmt.Sprintf(body, uri)
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
}
