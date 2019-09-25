package main

import (
	"math/big"
	"math/rand"
	"time"
)

// MillerRabinTest checks if provided number n is prime using k iterations of Miller Rabin test
// Return true if provided number is prime, and false if it's composite
func MillerRabinTest(n *big.Int, rounds int) bool {
	smallPrimes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	// Check if n is divisible by small prime numbers
	for _, prime := range smallPrimes {
		if big.NewInt(prime).Cmp(n) == 0 {
			continue
		}
		reminder := new(big.Int)
		reminder.Mod(n, big.NewInt(prime))
		if reminder.Cmp(big.NewInt(0)) == 0 {
			return false
		}
	}

	// Perform Miller-Rabin test r times
	for i := 0; i < rounds; i++ {
		d := new(big.Int)
		d.Sub(n, big.NewInt(1))
		// Find d, k such as d*(2**k) = n -1
		k, d := breaknumber(d)
		a := new(big.Int)
		r := getRandom()
		a.Rand(r, new(big.Int).Sub(n, big.NewInt(2)))
		a.Add(a, big.NewInt(2))
		x := new(big.Int).Exp(a, d, n)

		// Check a**d mod n == 1 or a**d mod n == -1
		if x.Cmp(big.NewInt(1)) == 0 || (x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0) {
			continue
		}

		for j := 0; j < k; j++ {
			// Check a**(2*d) mod n == -1 otherwise n is composite
			x = new(big.Int).Exp(x, big.NewInt(2), n)
			if x.Cmp(big.NewInt(1)) == 0 {
				return false
			}
			if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
				break
			}
		}
		if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
			return false
		}
	}
	return true
}

// LucasLehmetTest checks if Mesenne number M=2**p -1 is prime
// and return true, M if M is prime or false, 0 otherwise
func LucasLehmetTest(p *big.Int) (bool, *big.Int) {
	IsPrime := MillerRabinTest(p, 10)
	if !IsPrime {
		return false, big.NewInt(0)
	}
	bigTwo := big.NewInt(2)
	// Compute p-th Mersenne number Mp
	MeresenneNumber := new(big.Int).Sub(new(big.Int).Exp(bigTwo, p, nil), big.NewInt(1))
	residue := big.NewInt(4)
	k := big.NewInt(1)
	// Find (p-1)th Lucas number Lp+1
	for k.Cmp(new(big.Int).Sub(p, big.NewInt(1))) == -1 {
		residue = residue.Mod(residue.Sub(residue.Exp(residue, bigTwo, nil), bigTwo), MeresenneNumber)
		k = k.Add(k, big.NewInt(1))
	}
	// If Mp % Lp+1 = 0, then Mp is prime
	if residue.Cmp(big.NewInt(0)) == 0 {
		return true, MeresenneNumber
	}
	return false, big.NewInt(0)
}

func breaknumber(d *big.Int) (int, *big.Int) {
	// return k, s such as (2**k)*s = d
	k := 0
	reminder := new(big.Int)
	for (reminder.Mod(d, big.NewInt(2))).Cmp(big.NewInt(0)) == 0 {
		d.Div(d, big.NewInt(2))
		k++
	}
	return k, d
}

func getRandom() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r
}

// LucasTest check if provided number n is strong Lucas pseudoprime
func LucasTest(n *big.Int) bool {
	if isPerfectSquare(n) {
		return false
	}
	// Find correct D and P value for Lucas test
	var d int64 = 1
	var pInt int64 = 3
	for ; ; pInt++ {
		if pInt > 1000 {
			panic("Cannot find Jacobi(D/n) = -1 for " + n.String())
		}
		d = pInt*pInt - 4
		jacobi := big.Jacobi(big.NewInt(d), n)
		if jacobi == -1 {
			break
		}
	}

	// Set Lucas test P, Q parameters
	p := big.NewInt(pInt)
	q := big.NewInt(1)

	// find l,s such as l*(2**s) = n - jacob(d/n)
	tmp := new(big.Int).Add(n, big.NewInt(1))
	s, l := breaknumber(tmp)

	// U[1] = 1, V[1] = P
	uk, vk := big.NewInt(1), new(big.Int).Set(p)
	k := big.NewInt(1)
	for i := int(l.BitLen()) - 1; i > 0; i-- {
		// Compute U[l] and V[l]
		if l.Bit(i-1) == 0 {
			// k' = 2k
			// U[2k] = U[k] * V[k]
			// V[2k] = V[k] * V[k] - 2 * (Q**k)
			uk = uk.Mul(uk, vk)
			uk = uk.Mod(uk, n)

			vk = new(big.Int).Exp(vk, big.NewInt(2), n)
			sqrQ := new(big.Int).Exp(q, k, n)
			tmp := new(big.Int).Mul(sqrQ, big.NewInt(2))
			vk = new(big.Int).Sub(vk, tmp)
			vk = vk.Mod(vk, n)

			k = k.Mul(k, big.NewInt(2))
		} else {
			// k' = 2k+1
			// U[2k +1] = (P * U[2k] + V[2k]) / 2
			// V[2k +1] = (D * U[2k] + P * V[2k]) / 2
			uk = uk.Mul(uk, vk)
			uk = uk.Mod(uk, n)

			vk = new(big.Int).Exp(vk, big.NewInt(2), n)
			sqrQ := new(big.Int).Exp(q, k, n)
			tmp := new(big.Int).Mul(sqrQ, big.NewInt(2))
			vk = new(big.Int).Sub(vk, tmp)
			vk = vk.Mod(vk, n)

			tu := new(big.Int).Set(uk)
			tv := new(big.Int).Set(vk)

			uk = uk.Mul(uk, p)
			uk = uk.Add(uk, vk)
			if new(big.Int).Mod(uk, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
				uk = uk.Add(uk, n)
			}
			uk = uk.Div(uk, big.NewInt(2))
			uk = uk.Mod(uk, n)

			vk = new(big.Int).Mul(tu, big.NewInt(d))
			vk = new(big.Int).Add(vk, new(big.Int).Mul(p, tv))
			if new(big.Int).Mod(vk, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
				vk = vk.Add(vk, n)
			}
			vk = vk.Div(vk, big.NewInt(2))
			vk = vk.Mod(vk, n)

			k = k.Mul(k, big.NewInt(2))
			k = new(big.Int).Add(k, big.NewInt(1))
		}
	}

	// Check U[l] = 0 mod n
	if big.NewInt(0).Cmp(new(big.Int).Mod(uk, n)) == 0 {
		return true
	}

	// 	Check V[l] = 0 mod n
	if big.NewInt(0).Cmp(new(big.Int).Mod(vk, n)) == 0 {
		return true
	}
	for i := 1; i < s; i++ {
		// 	Check V[l * 2**i] = 0 mod n for each i < s
		vk = new(big.Int).Exp(vk, big.NewInt(2), n)
		sqrQ := new(big.Int).Exp(q, k, n)
		tmp := new(big.Int).Mul(sqrQ, big.NewInt(2))
		vk = new(big.Int).Sub(vk, tmp)
		vk = vk.Mod(vk, n)

		k = k.Mul(k, big.NewInt(2))

		if big.NewInt(0).Cmp(new(big.Int).Mod(vk, n)) == 0 {
			return true
		}
	}

	return false
}

func isPerfectSquare(n *big.Int) bool {
	// Return true if n is perfect square
	z := new(big.Int).Sqrt(n)
	z = z.Exp(z, big.NewInt(2), nil)
	if z.Cmp(n) == 0 {
		return true
	}
	return false
}
