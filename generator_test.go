package main

import (
	"testing"
)

func TestPrimeGenerator(t *testing.T) {
	for i := 0; i < 10; i++ {
		prime := GenerateRandomPrime(1024)
		if !prime.ProbablyPrime(10) {
			t.Errorf("GenerateRandomPrime returned %v, which is not prime number", prime)
		}
	}
}
