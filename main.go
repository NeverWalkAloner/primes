package main

import (
	"flag"
	"fmt"
	"math/big"
)

func main() {
	opt := flag.String("opt", "soe", "Options: soe, soa, mrt, lt, llt, gen")
	number := flag.String("number", "127", "Number to test primility")
	max := flag.Int("max", 100, "Max limit for sieves")
	size := flag.Int("size", 100, "Bit length of generated prime number")

	var primes []int

	flag.Parse()

	switch *opt {
	case "soe":
		primes = Eratosthenes(*max)
		fmt.Println(primes)
		fmt.Println(len(primes), "prime numbers generated")
	case "soa":
		primes = Atkin(*max)
		fmt.Println(primes)
		fmt.Println(len(primes), "prime numbers generated")
	case "mrt":
		n := new(big.Int)
		n, ok := n.SetString(*number, 10)
		if !ok {
			fmt.Println("cannot conver string value to BigInt")
		}
		result := MillerRabinTest(n, 5)
		fmt.Println("Is probable prime: ", result)
	case "lt":
		n := new(big.Int)
		n, ok := n.SetString(*number, 10)
		if !ok {
			fmt.Println("cannot conver string value to BigInt")
		}
		result := LucasTest(n)
		fmt.Println("Is probable prime: ", result)
	case "llt":
		n := new(big.Int)
		n, ok := n.SetString(*number, 10)
		if !ok {
			fmt.Println("cannot conver string value to BigInt")
		}
		result, M := LucasLehmetTest(n)
		fmt.Printf("Is mersenne number 2**%v=%v prime: %v\n", n, M, result)
	case "gen":
		prime := GenerateRandomPrime(*size)
		fmt.Println(prime)
	}
}
