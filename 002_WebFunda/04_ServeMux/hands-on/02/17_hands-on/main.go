package main

import (
	"bufio"
	"fmt"
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

	log.Printf("URL   : %s\n", uri)
	log.Printf("METHOD: %s\n", method)

	if uri == "/" && method == "GET" {
		home(c)
	} else if uri == "/apply" && method == "GET" {
		apply(c)
	} else if uri == "/apply" && method == "POST" {
		posted(c)
	} else {
		body := `
		<!DOCTYPE html><html lang="en"><head><meta charset="utf-8">
		<title>PAGE NOT FOUND</title></head><body>
		<h1>Could not find your Page</h1><a href="/">Home</a></body></html>
		`
		body = fmt.Sprintf(body, uri, method)
		io.WriteString(c, "HTTP/1.1 404 Not Found\r\n")
		fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
		fmt.Fprintf(c, "Content-Type: text/html\r\n")
		io.WriteString(c, "\r\n")
		fmt.Fprintf(c, body)
	}
}

func home(c net.Conn) {
	body := `
    <!DOCTYPE html><html lang="en"><head><meta charset="utf-8">
	<title>HOME PAGE</title></head><body>
	<h1> This is the Home Page </h1><a href="/apply">Application</a></body></html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	fmt.Fprintf(c, body)
}

func apply(c net.Conn) {
	body := `
    <!DOCTYPE html><html lang="en"><head><meta charset="utf-8">
	<title>Application PAGE</title></head><body>
	<h1> Form </h1><form method="POST">
	<input type="submit" value="Submit the Form">
	</form><a href="/">Home</a></body></html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	fmt.Fprintf(c, body)
}

func posted(c net.Conn) {
	body := `
    <!DOCTYPE html><html lang="en"><head><meta charset="utf-8">
	<title>CONFIRMATION PAGE</title></head><body>
	<h1> Your Application has been submitted successfully </h1>
	<a href="/">Home</a></body></html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	fmt.Fprintf(c, body)
}
