/*
Problem 41: Pandigital Primes
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once.
For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
*/

package problems

import (
	"github.com/zolrath/euler/util/primes"
	"github.com/zolrath/euler/util/sink"
)

const ANSWER_041 = 7652413

func Euler041() int {
	pl := primes.Sieve(7654321)
	largest := 0
	for _, p := range pl {
		if sink.IsPandigital(p) {
			largest = p
		}

	}
	return largest
}
