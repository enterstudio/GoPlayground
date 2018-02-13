package main

import (
	"bufio"
	"fmt"
	"html"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func handlehttp(conn net.Conn) {

	// Terminate the Connection at the end
	defer conn.Close()

	// Create the Scanner
	sn := bufio.NewScanner(conn)

	var protocol, host string
	header := make(map[string]string)
	i := 0
	// Walk through the Lines
	for sn.Scan() {

		line := sn.Text()

		// End of Processing Loop
		if line == "" {
			break
		}

		log.Println(line)

		switch i {
		case 0: // Protocol Process
			fd := strings.Fields(line)
			protocol = fd[0]
			host = fd[1]
		case 1:
			fd := strings.Fields(line)
			host = fd[1] + host
		default:
			// Processes the Headers
			sp := strings.SplitN(line, ":", 2)
			header[sp[0]] = sp[1]
		}

		i++
	}

	resp := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Protocol Processor</title>
</head>
<body>
	<hr>
	<h1>POST Request Form</h1>
	<form method="POST">
		<input type="text" name="key" value="">
		<input type="submit">
	</form>
	<hr>
	<h1>GET Request Form</h1>
	<a href="">GET Request</a>
	<hr>
`
	resp += "<p>Protocol : " + protocol + "<br> Host : <pre>" + html.EscapeString(host) + "</pre>"
	resp += "</p>"
	// Add Post Request data
	if protocol == "POST" || protocol == "PUT" {
		amt, _ := strconv.Atoi(header["Content-Length"])
		buf := make([]byte, amt)
		io.ReadFull(conn, buf)
		log.Println(" BODY: " + string(buf))
		// in buf we will have the POST content
		resp += "<pre> BODY: <br>" + html.EscapeString(string(buf)) + "</pre>"
	}
	resp += "<p> Headers: <br><pre>"
	for k, v := range header {
		resp += html.EscapeString(fmt.Sprintf("%20s : %s", k, v)) + "<br>"
	}
	resp += "</pre></p>"
	resp += "</body></html>"

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(resp))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, resp)
	log.Println(" Protocol : ", protocol)
	log.Println(" Headers : ", header)

}
func main() {

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer server.Close()

	log.Println(" Server started on localhost:8080 ")
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		go handlehttp(conn)
	}
}
