package main

import (
	"math/big"
	"testing"
)

func TestBreaknumber(t *testing.T) {
	// 272 = 16 * 17
	k, d := breaknumber(big.NewInt(272))

	if k != 4 {
		t.Errorf("Invalid value for k in breaknumber function expected 4, got %v", k)
	}

	if d.Cmp(big.NewInt(17)) != 0 {
		t.Errorf("Invalid value for d in breaknumber function expected 17, got %v", d)
	}
}

func TestMillerRabinWithFirstPrimes(t *testing.T) {
	arr := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83,
		89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181,
		191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283,
		293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523,
		541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647,
		653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773,
		787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911,
		919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997,
	}
	for _, a := range arr {
		n := big.NewInt(int64(a))

		result := MillerRabinTest(n, 10)

		if !result {
			t.Errorf("Invalid result in MillerRabinTest function for %v", n)
		}
	}
}

func TestMillerRabinWithPrimeNumber(t *testing.T) {
	primes := []string{
		"4611252981035380070313820144626939638615800115895757194765710672339892046994852722873028987577596159",
		"8672946699202715709509784660407977259465411153290803244410824992175935872689971595932039949736270607",
		"5221095160247394801617874835253681796697681379092384525204442267919561210219914160638328876842308927",
		"7653081782446593267761537482475681411664085331878780588228282279785116450762101690039122793930755481",
		"3672490010901911285015220763809528549615365199577573443743370459546164777918771859397904359897052377",
	}
	for _, prime := range primes {
		n := new(big.Int)
		n, ok := n.SetString(prime, 10)
		if !ok {
			t.Errorf("cannot conver string value to BigInt")
		}
		result := MillerRabinTest(n, 10)

		if !result {
			t.Errorf("Invalid result in MillerRabinTest function expected true, got %v", result)
		}
	}
}

func TestMillerRabinWithCompositeNumber(t *testing.T) {
	composites := []string{
		"3571252981035380070313820144626939638615800115895757194765710672339892046994852722873028987577596173",
		"9852946699202715709509784660407977259465411153290803244410824992175935872689971595932039949736270231",
		"9871095160247394801617874835253681796697681379092384525204442267919561210219914160638328876842308987",
		"6543081782446593267761537482475681411664085331878780588228282279785116450762101690039122793930755197",
		"1112490010901911285015220763809528549615365199577573443743370459546164777918771859397904359897052097",
	}
	for _, composite := range composites {
		n := new(big.Int)
		n, ok := n.SetString(composite, 10)
		if !ok {
			t.Errorf("cannot conver string value to BigInt")
		}
		result := MillerRabinTest(n, 10)

		if result {
			t.Errorf("Invalid result in MillerRabinTest function expected false, got %v", result)
		}
	}
}

func TestMillerRabinWithCarmichaelNumber(t *testing.T) {
	carmichaels := []string{
		"1590231231043178376951698401",
		"6553130926752006031481761",
		"87674969936234821377601",
	}
	for _, carmichael := range carmichaels {
		n := new(big.Int)
		n, ok := n.SetString(carmichael, 10)
		if !ok {
			t.Errorf("cannot conver string value to BigInt")
		}
		result := MillerRabinTest(n, 10)

		if result {
			t.Errorf("Invalid result in MillerRabinTest function expecting false, got %v", result)
		}
	}
}

func TestLicasLehmerWithPrimeNumber(t *testing.T) {
	primes := []int{7, 17, 127, 521, 607}
	MersennePrimes := []string{
		"127",
		"131071",
		"170141183460469231731687303715884105727",
		"6864797660130609714981900799081393217269435300143305409394463459185543183397656052122559640661454554977296311391480858037121987999716643812574028291115057151",
		"531137992816767098689588206552468627329593117727031923199444138200403559860852242739162502265229285668889329486246501015346579337652707239409519978766587351943831270835393219031728127",
	}
	for i, prime := range primes {
		p := big.NewInt(int64(prime))
		isPrime, M := LucasLehmetTest(p)

		if !isPrime {
			t.Errorf("Invalid result in MillerRabinTest function expected true, got %v", isPrime)
		}

		n := new(big.Int)
		n, ok := n.SetString(MersennePrimes[i], 10)
		if !ok {
			t.Errorf("cannot conver string value to BigInt")
		}

		if n.Cmp(M) != 0 {
			t.Errorf("Invalid value of Mersenne prime expected %v, got %v", n, M)
		}
	}
}

func TestLicasLehmerWithCopositeNumber(t *testing.T) {
	invalidPrimes := []int{11, 29, 149, 523, 617}

	for _, prime := range invalidPrimes {
		p := big.NewInt(int64(prime))
		isPrime, _ := LucasLehmetTest(p)

		if isPrime {
			t.Errorf("Invalid result in LicasLehmerTest function expected false, got %v", isPrime)
		}
	}
}

func TestLucasWithFirstPrimes(t *testing.T) {
	arr := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83,
		89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181,
		191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283,
		293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523,
		541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647,
		653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773,
		787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911,
		919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997,
	}
	for _, a := range arr {
		n := big.NewInt(int64(a))

		result := LucasTest(n)

		if !result {
			t.Errorf("Invalid result in LucasTest for %v", n)
		}
	}
}

func TestLucasWithPrimeNumber(t *testing.T) {
	primes := []string{
		"4611252981035380070313820144626939638615800115895757194765710672339892046994852722873028987577596159",
		"8672946699202715709509784660407977259465411153290803244410824992175935872689971595932039949736270607",
		"5221095160247394801617874835253681796697681379092384525204442267919561210219914160638328876842308927",
		"7653081782446593267761537482475681411664085331878780588228282279785116450762101690039122793930755481",
		"3672490010901911285015220763809528549615365199577573443743370459546164777918771859397904359897052377",
	}
	for _, prime := range primes {
		n := new(big.Int)
		n, ok := n.SetString(prime, 10)
		if !ok {
			t.Errorf("cannot conver string value to BigInt")
		}
		result := LucasTest(n)

		if !result {
			t.Errorf("Invalid result in LucasTest function expected true, got %v", result)
		}
	}
}

func TestLucasWithCompositeNumber(t *testing.T) {
	composites := []string{
		"3571252981035380070313820144626939638615800115895757194765710672339892046994852722873028987577596173",
		"9852946699202715709509784660407977259465411153290803244410824992175935872689971595932039949736270231",
		"9871095160247394801617874835253681796697681379092384525204442267919561210219914160638328876842308987",
		"6543081782446593267761537482475681411664085331878780588228282279785116450762101690039122793930755197",
		"1112490010901911285015220763809528549615365199577573443743370459546164777918771859397904359897052097",
	}
	for _, composite := range composites {
		n := new(big.Int)
		n, ok := n.SetString(composite, 10)
		if !ok {
			t.Errorf("cannot conver string value to BigInt")
		}
		result := LucasTest(n)

		if result {
			t.Errorf("Invalid result in LucasTest function expected false, got %v", result)
		}
	}
}
