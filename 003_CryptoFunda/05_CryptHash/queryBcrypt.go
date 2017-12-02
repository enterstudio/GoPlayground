package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "TestingPassword"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	fmt.Println(" 1st #Generated: ", base64.RawStdEncoding.EncodeToString(hash))
	hash, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	fmt.Println(" 2nd #Generated: ", base64.RawStdEncoding.EncodeToString(hash))
}
