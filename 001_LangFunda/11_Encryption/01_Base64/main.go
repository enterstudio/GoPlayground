package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	tst := "Hello"
	s := base64.StdEncoding.EncodeToString([]byte(tst))
	fmt.Println("Test String:", tst)
	fmt.Println("Test String:", []byte(tst))
	fmt.Println("Encoded:", s)
	fmt.Println("Encoded:", []byte(s))
	d, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalln("Decoding failed")
	}
	fmt.Println("Decoded:", d)
	fmt.Println("Decoded:", string(d))
}
