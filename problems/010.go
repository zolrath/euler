/*
Problem 10: Summation of Primes
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
Answer: 142913828922
*/

package problems

import "math"

func simpleSieveSum(limit int) int {
	sieve := make([]int, limit+1)
	sum := 0
	for i := 2; i <= int(math.Sqrt(float64(limit))); i++ {
		if sieve[i] == 0 {
			for j := i * i; j <= limit; j += i {
				sieve[j] = 1
			}
		}
	}
	for i := 2; i <= limit; i++ {
		if sieve[i] == 0 {
			sum += i
		}
	}
	return sum
}

func Euler010() int {
	return simpleSieveSum(2000000)
}
