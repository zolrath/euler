// Package primes contains various methods of generating primes and prime related functions.
package primes

import "math"

// Nth returns the nth prime number.
func Nth(n int) int {
	// start with 2 in the primelist so we can only search odd numbers.
	primes := []int{2}
Primesearch:
	for i := 3; len(primes) < n; i += 2 {
		sqrt := int(math.Sqrt(float64(i)))
		for _, v := range primes {
			switch {
			case v > sqrt:
				break
			case i%v == 0:
				continue Primesearch
			}
		}
		primes = append(primes, i)
	}
	return primes[n-1]
}

// N returns a prime list with N primes.
func N(n int) []int {
	// start with 2 in the primelist so we can only search odd numbers.
	primes := []int{2}
Primesearch:
	for i := 3; len(primes) < n; i += 2 {
		sqrt := int(math.Sqrt(float64(i)))
		for _, v := range primes {
			switch {
			case v > sqrt:
				break
			case i%v == 0:
				continue Primesearch
			}
		}
		primes = append(primes, i)
	}
	return primes
}

// Sieve uses a Sieve of Eratosthenes, generating the prime numbers less than the limit.
func Sieve(limit int) []int {
	sieve := make([]int, limit+1)
	primes := []int{2}
	for i := 3; i <= int(math.Sqrt(float64(limit))); i += 2 {
		if sieve[i] == 0 {
			for j := i * i; j <= limit; j += i {
				sieve[j] = 1
			}
		}
	}
	for i := 3; i <= limit; i += 2 {
		if sieve[i] == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}

// SieveMap uses a Sieve of Eratosthenes, generating the prime numbers less than the limit.
// It returns a map for fast isPrime checking.
func SieveMap(limit int) map[int]bool {
	sieve := make([]int, limit+1)
	primes := map[int]bool{2: true}
	for i := 3; i <= int(math.Sqrt(float64(limit))); i += 2 {
		if sieve[i] == 0 {
			for j := i * i; j <= limit; j += i {
				sieve[j] = 1
			}
		}
	}
	for i := 3; i <= limit; i += 2 {
		if sieve[i] == 0 {
			primes[i] = true
		}
	}
	return primes
}

// IsPrime uses a SieveMap to determine if a number is prime or not.
func IsPrime(n int, primeMap map[int]bool) bool {
	if primeMap == nil {
		sieve := Sieve(n + 1)
		if sieve[len(sieve)-1] == n {
			return true
		}
		return false
	}
	if _, ok := primeMap[n]; !ok {
		return false
	}
	return true
}

func AllArePrime(nums []int, primeMap map[int]bool) bool {
	for _, v := range nums {
		if !IsPrime(v, primeMap) {
			return false
		}
	}
	return true
}

// PrimeFactors turns a map of the prime factors of n.
// Each key represents a prime factor.
// Each value represents the number of times the prime factors into n.
func PrimeFactors(n int) map[int]int {
	pfs := map[int]int{}
	max := int(math.Sqrt(float64(n)))
	for i := 2; i < max; i++ {
		for n%i == 0 {
			n = n / i
			pfs[i] = pfs[i] + 1
		}
	}
	if n > 1 {
		pfs[n] = 1
	}
	return pfs
}
