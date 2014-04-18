/*
Problem 10: Summation of Primes
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package problems

import "github.com/zolrath/euler/util/primes"

const ANSWER_010 = 142913828922

func Euler010() int {
	sum := 0
	p := primes.Sieve(2000000)
	for _, v := range p {
		sum += v
	}
	return sum
}
