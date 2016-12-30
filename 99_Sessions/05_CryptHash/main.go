package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	password := "hello"

	hash, salt, _ := ScryptHash(password)
	fmt.Println(" Scrypt Hash")
	fmt.Println(base64.StdEncoding.EncodeToString(hash))
	fmt.Println(base64.StdEncoding.EncodeToString(salt))
	hash, salt, _ = PBKDF2Hash(password)
	fmt.Println(" PBKDF2 Hash")
	fmt.Println(base64.StdEncoding.EncodeToString(hash))
	fmt.Println(base64.StdEncoding.EncodeToString(salt))
	hash, salt, _ = BcryptHash(password)
	fmt.Println(" BCrypt Hash")
	fmt.Println(base64.StdEncoding.EncodeToString(hash))
	fmt.Println(base64.StdEncoding.EncodeToString(salt))
}
