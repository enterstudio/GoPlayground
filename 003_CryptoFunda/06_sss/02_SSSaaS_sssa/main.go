package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SSSaaS/sssa-golang"
)

func main() {
	secret := "Hari Aum Tat Sat!"
	num_shares := 5
	min_reconstruct := 2
	fmt.Printf("\n\nThe Secret is %q\n\n", secret)
	fmt.Printf(" We would be splitting this into %d number of shares\n", num_shares)
	fmt.Printf(" Also a minimum %d number of shares are need to reconstruct the secret\n", min_reconstruct)
	fmt.Println("\n ...Hari Aum!...")

	shares, err := sssa.Create(min_reconstruct, num_shares, secret)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n Here are the Shares:\n\n")
	for i, k := range shares {
		fmt.Printf(" - Share %d = %s\n", i, k)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n1 := r.Intn(num_shares)
	n2 := r.Intn(num_shares)
	for n1 == n2 {
		n1 = r.Intn(num_shares)
		n2 = r.Intn(num_shares)
	}

	fmt.Printf("\n Random Selected Pairs: %d , %d\n\n", n1, n2)

	p1 := shares[n1]
	p2 := shares[n2]
	fmt.Printf("\n The Pairs Are :\n")
	fmt.Printf("\n P1 [%d] = %q Valid = %v\n", n1, p1, sssa.IsValidShare(p1))
	fmt.Printf("\n P2 [%d] = %q Valid = %v\n", n2, p2, sssa.IsValidShare(p2))

	recomb := []string{p1, p2}
	rec, err := sssa.Combine(recomb)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nRecovered String Is: %q\n\n\n", rec)
}
