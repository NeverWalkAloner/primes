package main

import "testing"

var firstPrimes = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
	59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181,
	191, 193, 197, 199,
}

func TestEratosthenes(t *testing.T) {
	generatedPrimes := Eratosthenes(200)

	if len(generatedPrimes) != len(firstPrimes) {
		t.Errorf("Invalid count of primes generated")
	}

	for i, prime := range generatedPrimes {
		if prime != firstPrimes[i] {
			t.Errorf("%d-th number incorrect, got: %d, want: %d.", i, prime, firstPrimes[i])
		}
	}
}

func TestAtkin(t *testing.T) {
	generatedPrimes := Atkin(200)

	if len(generatedPrimes) != len(firstPrimes) {
		t.Errorf("Invalid count of primes generated")
	}

	for i, prime := range generatedPrimes {
		if prime != firstPrimes[i] {
			t.Errorf("%d-th number incorrect, got: %d, want: %d.", i, prime, firstPrimes[i])
		}
	}
}
