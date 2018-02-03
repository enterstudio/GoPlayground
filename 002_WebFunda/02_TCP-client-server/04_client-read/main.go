package main

import (
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(bs))
}
