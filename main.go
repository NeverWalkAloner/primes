package main

import (
	"flag"
	"fmt"
	"math/big"
	"time"
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
		start := time.Now()
		primes = Eratosthenes(*max)
		elapsed := time.Since(start)
		fmt.Println(primes)
		fmt.Println("Eratosthenes sieve took", elapsed)
		fmt.Println(len(primes), "prime numbers generated")
	case "soa":
		start := time.Now()
		primes = Atkin(*max)
		elapsed := time.Since(start)
		fmt.Println(primes)
		fmt.Println("Atkin sieve took", elapsed)
		fmt.Println(len(primes), "prime numbers generated")
	case "mrt":
		start := time.Now()
		n := new(big.Int)
		n, ok := n.SetString(*number, 10)
		if !ok {
			fmt.Println("cannot conver string value to BigInt")
		}
		result := MillerRabinTest(n, 5)
		elapsed := time.Since(start)
		fmt.Println("Is probable prime: ", result)
		fmt.Println("Miller-Rabin test took", elapsed)
	case "lt":
		start := time.Now()
		n := new(big.Int)
		n, ok := n.SetString(*number, 10)
		if !ok {
			fmt.Println("cannot conver string value to BigInt")
		}
		result := LucasTest(n)
		elapsed := time.Since(start)
		fmt.Println("Is probable prime: ", result)
		fmt.Println("Lucas test took", elapsed)
	case "llt":
		start := time.Now()
		n := new(big.Int)
		n, ok := n.SetString(*number, 10)
		if !ok {
			fmt.Println("cannot conver string value to BigInt")
		}
		result, M := LucasLehmetTest(n)
		elapsed := time.Since(start)
		fmt.Printf("Is mersenne number 2**%v=%v prime: %v\n", n, M, result)
		fmt.Println("Licas-Lehmer test took", elapsed)
	case "gen":
		start := time.Now()
		prime := GenerateRandomPrime(*size)
		fmt.Println(prime)
		elapsed := time.Since(start)
		fmt.Println("Prime number generation took", elapsed)
	}
}
