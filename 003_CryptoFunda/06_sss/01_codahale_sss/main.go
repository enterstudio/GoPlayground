package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/codahale/sss"
)

func genSelector(max byte) []byte {
	b := make([]byte, 2)
	rand.Read(b) // First Attempt
	for i, _ := range b {
		b[i] = (b[i] % max) + 1
	}
	return b
}

func main() {
	secret := "Hari Aum Tat Sat!"
	num_shares := byte(5)
	min_reconstruct := byte(2)
	fmt.Printf("\n\nThe Secret is %q\n\n", secret)
	fmt.Printf(" We would be splitting this into %d number of shares\n", num_shares)
	fmt.Printf(" Also a minimum %d number of shares are need to reconstruct the secret\n", min_reconstruct)
	fmt.Println("\n ..Hari Aum!..")

	shares, err := sss.Split(num_shares, min_reconstruct, []byte(secret))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n Here are the Shares:\n\n")

	// Printing the Shares in order
	for i, k := range shares {
		fmt.Printf("Share - %d = %s - %s\n", i, base64.RawStdEncoding.EncodeToString(k), hex.EncodeToString(k))
	}
	b := genSelector(num_shares)
	for b[0] == b[1] { // Check they should not be equal
		b = genSelector(num_shares)
	}
	fmt.Printf("\n Random shares Chosen for Reconstruction: %v\n", b)

	// Note: Here share number is also important for Decoding the Calculation
	subset := map[byte][]byte{b[0]: shares[b[0]], b[1]: shares[b[1]]}
	fmt.Printf("\n Here are the Selected Shares:\n\n")
	for i, k := range subset {
		fmt.Printf("Share - %d = %s - %s\n", i, base64.RawStdEncoding.EncodeToString(k), hex.EncodeToString(k))
	}

	recovered := string(sss.Combine(subset))
	fmt.Printf("\n Recovered secret is %q\n\n", recovered)
}
