package main

import (
	"bufio"
	"log"
	"net"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()
	var wg sync.WaitGroup
	log.Println(" Server Started at 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		wg.Add(1)
		go handle(conn, wg)
	}
	wg.Wait()
}

func handle(conn net.Conn, wg sync.WaitGroup) {
	defer conn.Close()
	scan := bufio.NewScanner(conn)
	data := ""
	for scan.Scan() {
		line := scan.Text()
		log.Println(line)
		data += line
	}

	// Does Not Reach Here Due to Scan Line Running
	log.Println("Done with the Connection")
	conn.Write([]byte("Done"))
	wg.Done()
}
