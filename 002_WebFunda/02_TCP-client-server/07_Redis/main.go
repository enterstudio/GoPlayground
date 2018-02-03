package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func handleRedis(conn net.Conn) {
	// Close connection at the end
	defer conn.Close()
	use := `
In Memory Database

	USE:
	SET <key> <value>
	GET <key>
	DEL <key>

	Example:
	SET fav chocolate
	GET fav				


`

	// Define Redis store
	store := make(map[string]string)

	// Print the Usage
	io.WriteString(conn, use)

	// Create the Scanner
	scan := bufio.NewScanner(conn)

	for scan.Scan() {
		line := scan.Text()

		// Dont Process Empty Lines
		if len(line) == 0 {
			continue
		}

		// Break it into fields
		fld := strings.Fields(line)
		log.Println(fld)
		// Process the Field
		switch fld[0] {
		case "GET":
			if len(fld) > 1 {
				if val, present := store[fld[1]]; present {
					io.WriteString(conn, val)
				}
			}
		case "SET":
			if len(fld) > 2 {
				key := fld[1]
				store[key] = fld[2]
			}
		case "DEL":
			if len(fld) > 1 {
				delete(store, fld[1])
			}
		default:
			io.WriteString(conn, use)
		}
		// Finally a Return in the endl
		io.WriteString(conn, "\r\n")
	}
}

func main() {

	server, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(" Starting Server localhost:9000")

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		go handleRedis(conn)
	}
}
