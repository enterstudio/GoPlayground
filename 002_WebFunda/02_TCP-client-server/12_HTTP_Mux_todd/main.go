/*

This program is specifically shaped to mimic the example shown
by Todd Mcleod for the HTTP Mux

Special Mention of the Particular Video:
https://www.udemy.com/go-programming-language/learn/v4/t/lecture/6027314?start=0

This is Section 3: Creating your own server
Video 28. TCP server - HTTP multiplexer

*/
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
		log.Fatalln(err)
	}
	log.Println(" Starting TCP Server at 8080")
	for {
		conn, err := svr.Accept()
		if err != nil {
			log.Fatalln(err)
		} else {
			go handle(conn)
		}
	}
}

func handle(conn net.Conn) {
	// For POST Request We need to make sure we get all the lines
	//  Generally the the Last line which has data comes after
	//  the connection timeout is reached
	err := conn.SetDeadline(time.Now().Add(200 * time.Millisecond))
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	request(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	line := 0
	ln := ""
	method := ""
	for scanner.Scan() {
		ln = scanner.Text()
		log.Println(ln)
		if line == 0 && len(ln) > 10 {
			method = strings.Fields(ln)[0]
			mux(conn, ln)
		}
		line++
	}
	if method == "POST" {
		log.Println(" POST Request Data:", ln)
	}
}

func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	if method == "GET" && uri == "/" {
		home(conn)
	} else if method == "GET" && uri == "/about" {
		about(conn)
	} else if method == "GET" && uri == "/contact" {
		contact(conn)
	} else if method == "GET" && uri == "/apply" {
		form(conn)
	} else if method == "POST" && uri == "/apply" {
		processForm(conn)
	} else {
		page404(conn, uri)
	}
}

func printhttp(conn net.Conn, buf string, isOk bool) {
	if isOk {
		fmt.Fprintf(conn, "HTTP/1.1 200 Ok\r\n")
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n")
	}
	fmt.Fprintf(conn, "Date: %s\r\n", time.Now().String())
	fmt.Fprintf(conn, "Content-Size: %d\r\n", len(buf))
	fmt.Fprintf(conn, "Server: %s\r\n", "Golang Server")
	fmt.Fprintf(conn, "Content-Type: %s\r\n", "text/html")
	fmt.Fprintf(conn, "Connection: close\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "%s", buf)
}

func home(conn net.Conn) {
	page := `
	<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
	<html><head><title>Home Page</title></head>
	<style>.menu{display:flex;list-style: none;}.menu li{padding:1rem;}</style>
	<body>
	<ul class="menu">
	<li><a href="/">Home</a></li>
	<li><a href="/about">About</a></li>
	<li><a href="/contact">Contact</a></li>
	<li><a href="/apply">Apply</a></li>
	</ul>
	<h1>Welcome</h1>
	<p>This is the Home Page</p>
	</body></html>
	`
	printhttp(conn, page, true)
}

func about(conn net.Conn) {
	page := `
	<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
	<html><head><title>About Page</title></head>
	<style>.menu{display:flex;list-style: none;}.menu li{padding:1rem;}</style>
	<body>
	<ul class="menu">
	<li><a href="/">Home</a></li>
	<li><a href="/about">About</a></li>
	<li><a href="/contact">Contact</a></li>
	<li><a href="/apply">Apply</a></li>
	</ul>
	<h1>About Us</h1>
	<p>To know more is a noble endeavor</p>
	</body></html>	
	`
	printhttp(conn, page, true)
}

func contact(conn net.Conn) {
	page := `
	<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
	<html><head><title>Contact Page</title></head>
	<style>.menu{display:flex;list-style: none;}.menu li{padding:1rem;}</style>
	<body>
	<ul class="menu">
	<li><a href="/">Home</a></li>
	<li><a href="/about">About</a></li>
	<li><a href="/contact">Contact</a></li>
	<li><a href="/apply">Apply</a></li>
	</ul>
	<h1>Contact Us</h1>
	<p>Contact Info ... info@example.com <br> call us at: +91-1234567890</p>
	</body></html>
	`
	printhttp(conn, page, true)
}

func form(conn net.Conn) {
	page := `
	<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
	<html><head><title>Apply Form</title></head>
	<style>.menu{display:flex;list-style: none;}.menu li{padding:1rem;}input, label{margin:8px;}</style>
	<body>
	<ul class="menu">
	<li><a href="/">Home</a></li>
	<li><a href="/about">About</a></li>
	<li><a href="/contact">Contact</a></li>
	<li><a href="/apply">Apply</a></li>
	</ul>
	<h1>Application Form</h1><p>Enter the Details:</p>
	<form method="POST"><label for="name">Name:</label><input type="text" name="name" id="name">
	<br><input type="submit" value="Submit Form"></form>
	</body></html>
	`
	printhttp(conn, page, true)
}

func processForm(conn net.Conn) {
	page := `
	<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
	<html><head><title>Successful Submission</title></head>
	<style>.menu{display:flex;list-style: none;}.menu li{padding:1rem;}</style>
	<body>
	<ul class="menu">
	<li><a href="/">Home</a></li>
	<li><a href="/about">About</a></li>
	<li><a href="/contact">Contact</a></li>
	<li><a href="/apply">Apply</a></li>
	</ul>
	<h1>Application Submitted</h1>
	<p>We have received your application successfully</p>
	</body></html>
	`
	printhttp(conn, page, true)
}

func page404(conn net.Conn, uri string) {
	page := `
	<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
	<html><head><title>404 Not Found</title></head>
	<body>
	<h1>Not Found</h1>
	<p>The requested URL %s was not found on this server.</p>
	</body></html>
	`
	page = fmt.Sprintf(page, uri)
	printhttp(conn, page, false)
}
