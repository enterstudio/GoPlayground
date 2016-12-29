package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha256"
	"crypto/rand"
)

func main(){
	password := "Aum Sri Ganeshay Namh"
	fmt.Println(" Password: ", password)
	fmt.Println(" Password Bytes array: ", []byte(password))
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(" Error in Bcrypt Password Hashing ")
	} else {
		fmt.Println(" BCrypt Hash Byte Array:")
		fmt.Println(hash)
	}
	fmt.Println(" Standard PBKDF2 with \"Salt\", 512 iterations, 64bytes, Sha256 hashing")
	hash2 := pbkdf2.Key([]byte(password), []byte("Salt"), 512, 64, sha256.New)
	fmt.Println(hash2)
	fmt.Println(" Length of PBKDF2 Hash: ", len(hash2))
	salt := make([]byte, 64)
	_, err = rand.Read(salt)
	if err!= nil {
		fmt.Println(" Error in Random Salt Generation")
	}
	fmt.Println(" Random Salt for PBKDF2 Byte Array: ", salt)
	hash2 =  pbkdf2.Key([]byte(password), salt, 512, 64, sha256.New)
	fmt.Println(" Standard PBKDF2 with CSPRNG Salt, 512 iterations, 64bytes, Sha256 hashing")
	fmt.Println(hash2)
}
