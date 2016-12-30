package main

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

const (
	saltLength       = 64
	scrypt_N         = 16384
	scrypt_r         = 8
	scrypt_p         = 1
	scrypt_KeyLength = 64
	pbkdf2_KeyLength = 64
	pbkdf2_Iteration = 512
)

func generateRandomSalt(length uint) ([]byte, error) {
	// Generate the Salt
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	return salt, err
	//return []byte("Salt"), nil
}

func BcryptHash(password string) ([]byte, []byte, error) {

	// Get Password Length
	lenpass := len(password)
	if lenpass == 0 {
		return []byte{}, []byte{}, errors.New("Empty Password String")
	}

	// Generate the Salt
	salt, err := generateRandomSalt(saltLength)
	if err != nil {
		return []byte{}, []byte{}, errors.New(" Unable to generate Salt " + err.Error())
	}

	// Generate the Hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password+string(salt)), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, []byte{}, errors.New(" Unable to generate Bcrypt Hash " + err.Error())
	}

	return hash, salt, nil
}

func ScryptHash(password string) ([]byte, []byte, error) {

	// Get Password Length
	lenpass := len(password)
	if lenpass == 0 {
		return []byte{}, []byte{}, errors.New("Empty Password String")
	}

	// Generate the Salt
	salt, err := generateRandomSalt(saltLength)
	if err != nil {
		return []byte{}, []byte{}, errors.New(" Unable to generate Salt " + err.Error())
	}

	// Generate the Hash
	hash, err := scrypt.Key([]byte(password), salt, scrypt_N, scrypt_r, scrypt_p, scrypt_KeyLength)
	if err != nil {
		return []byte{}, []byte{}, errors.New(" Unable to generate Scrypt Hash " + err.Error())
	}

	return hash, salt, nil
}

func PBKDF2Hash(password string) ([]byte, []byte, error) {

	// Get Password Length
	lenpass := len(password)
	if lenpass == 0 {
		return []byte{}, []byte{}, errors.New("Empty Password String")
	}

	// Generate the Salt
	salt, err := generateRandomSalt(saltLength)
	if err != nil {
		return []byte{}, []byte{}, errors.New(" Unable to generate Salt " + err.Error())
	}

	// Generate the Hash
	hash := pbkdf2.Key([]byte(password), salt, pbkdf2_Iteration, pbkdf2_KeyLength, sha256.New)

	return hash, salt, nil
}
