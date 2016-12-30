package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha256"
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

func main(){


	password := "Aum Sri Ganeshay Namh"
	fmt.Println("----------------")
	fmt.Println(" Password: ", password)
	fmt.Println(" Password Bytes array: ", []byte(password))
	fmt.Println("----------------")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(" Error in Bcrypt Password Hashing ")
	} else {
		fmt.Println(" BCrypt Hash Byte Array:")
		fmt.Println(hash)
		fmt.Println(" BCrypt Hash length: ", len(hash))
	}
	fmt.Println("----------------")
	/*
		Scrypt Parameters
	  N The CPU difficulty (must be a power of 2, > 1)
		r The memory difficulty
		p The parallel difficulty
	 */
	hash1, err := scrypt.Key([]byte(password), []byte("Salt"),16384,8,2,64)
	if err != nil {
		fmt.Println("Error Scrypt Hashing Failed ")
	} else {
		fmt.Println(" Standard Scrypt# with \"Salt\", N = 16384 r = 8 p = 1 for 64 Bytes")
		fmt.Println(hash1)
	}
	fmt.Println("----------------")
	fmt.Println(" Standard PBKDF2 with \"Salt\", 512 iterations, 64bytes, Sha256 hashing")
	hash2 := pbkdf2.Key([]byte(password), []byte("Salt"), 512, 64, sha256.New)
	fmt.Println(hash2)
	fmt.Println("----------------")
	salt := make([]byte, 64)
	_, err = rand.Read(salt)
	if err!= nil {
		fmt.Println(" Error in Random Salt Generation")
	}
	fmt.Println(" Random Salt 64Byte Array: ", salt)
	fmt.Println("----------------")
	hash1, err = scrypt.Key([]byte(password), salt, 16384,8,2,64)
	if err != nil {
		fmt.Println("Error Scrypt Hashing Failed ")
	} else {
		fmt.Println(" Standard Scrypt# with Random, N = 16384 r = 8 p = 1 for 64 Bytes")
		fmt.Println(hash1)
	}
	fmt.Println("----------------")
	hash2 =  pbkdf2.Key([]byte(password), salt, 512, 64, sha256.New)
	fmt.Println(" Standard PBKDF2 with CSPRNG Salt, 512 iterations, 64bytes, Sha256 hashing")
	fmt.Println(hash2)
	fmt.Println("----------------")
}
