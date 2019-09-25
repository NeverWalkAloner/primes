package main

import (
	"crypto/rand"
	"math/big"
)

// GenerateRandomPrime returns random prime number which pass Baillie–PSW primality test
func GenerateRandomPrime(l int) *big.Int {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(int64(l)), nil).Sub(max, big.NewInt(1))

	//Generate random in range [0, max)
	for {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic("Cannot generate bigInt")
		}
		if MillerRabinTest(n, 10) && LucasTest(n) {
			return n
		}
	}
}
