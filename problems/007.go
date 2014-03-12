/*
Problem 7: 10001st Prime
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10,001st prime number?
Answer: 104743
*/

package problems

import "math"

func nthPrime(n int) int {
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

func Euler007() int {
	return nthPrime(10001)
}
