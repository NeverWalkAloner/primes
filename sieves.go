package main

import (
	"math"
)

// Eratosthenes returns list of all prime numbers less than n, Eratosthenes sieve is using
func Eratosthenes(n int) []int {
	arr := make([]bool, n)
	for i := 2; i <= int(math.Sqrt(float64(n)+1)); i++ {
		if arr[i] == false {
			for j := i * i; j < n; j += i {
				arr[j] = true
			}
		}
	}
	var primes []int

	for i, isComposite := range arr {
		if i > 1 && !isComposite {
			primes = append(primes, i)
		}
	}

	return primes
}

// Atkin returns list of all prime numbers less than n, Atkin sieve is using
func Atkin(limit int) []int {
	x2, y2, n := 0, 0, 0
	sqrtLim := int(math.Sqrt(float64(limit)))
	arr := make([]bool, limit+1)
	arr[2], arr[3] = true, true

	for i := 1; i <= sqrtLim; i++ {
		x2 = i * i
		for j := 1; j <= sqrtLim; j++ {
			y2 = j * j
			n = 4*x2 + y2

			if (n <= limit) && (n%12 == 1 || n%12 == 5) {
				arr[n] = !arr[n]
			}

			n -= x2
			if (n <= limit) && (n%12 == 7) {
				arr[n] = !arr[n]
			}

			n -= 2 * y2
			if (i > j) && (n <= limit) && (n%12 == 11) {
				arr[n] = !arr[n]
			}
		}
	}

	for i := 5; i <= sqrtLim; i++ {
		if arr[i] {
			n = i * i
			for j := n; j <= limit; j += n {
				arr[j] = false
			}
		}
	}

	var primes []int
	for i, isPrime := range arr {
		if i > 1 && isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
