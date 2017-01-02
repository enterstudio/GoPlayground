package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	password := "hello"

	hash, salt, _ := ScryptHash(password)
	fmt.Println(" Scrypt Hash")
	fmt.Println(base64.RawStdEncoding.EncodeToString(hash))
	fmt.Println(base64.RawStdEncoding.EncodeToString(salt))
	combo := []byte(string(salt) + string(hash))
	if err := VerifyScrypt(password, combo); err == nil {
		fmt.Println(" Verified Scrypt Hash is valid")
	} else {
		fmt.Println(" Scrypt# is invalid " + err.Error())
	}
	hash, salt, _ = PBKDF2Hash(password)
	fmt.Println(" PBKDF2 Hash")
	fmt.Println(base64.RawStdEncoding.EncodeToString(hash))
	fmt.Println(base64.RawStdEncoding.EncodeToString(salt))
	combo = []byte(string(salt) + string(hash))
	if err := VerifyPBKDF2(password, combo); err == nil {
		fmt.Println(" Verified PBKDF2 Hash is valid")
	} else {
		fmt.Println(" PBKDF2# is invalid " + err.Error())
	}
	hash, salt, _ = BcryptHash(password)
	fmt.Println(" BCrypt Hash")
	fmt.Println(base64.RawStdEncoding.EncodeToString(hash))
	fmt.Println(base64.RawStdEncoding.EncodeToString(salt))

	/* // Need to Verify
	combo = []byte(string(salt) + string(hash))
	if err := VerifyBcrypt(password, combo); err == nil {
		fmt.Println(" Verified BCrypt Hash is valid")
	} else {
		fmt.Println(" BCrypt# is invalid " + err.Error())
	}*/
}
