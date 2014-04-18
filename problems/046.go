/*
Problem 46: Goldbach's Other Conjecture
It was proposed by Christian Goldbach that every odd composite number
can be written as the sum of a prime and twice a square.
9 = 7 + 2×1^2
15 = 7 + 2×2^2
21 = 3 + 2×3^2
25 = 7 + 2×3^2
27 = 19 + 2×2^2
33 = 31 + 2×1^2
It turns out that the conjecture was false.
What is the smallest odd composite that cannot be written
as the sum of a prime and twice a square?
*/

package problems

import (
	"github.com/zolrath/euler/util/primes"
	"github.com/zolrath/euler/util/sink"
)

const ANSWER_046 = 5777

func isTwiceASquare(n int) bool {
	return sink.IsPerfectSquare(n / 2)
}

func conjectureTrue(n int, primeList []int) bool {
	for i := 0; primeList[i] < n; i++ {
		p := primeList[i]
		rest := n - p
		if isTwiceASquare(rest) {
			return true
		}
	}
	return false
}

func Euler046() int {
	limit := 6000
	pl := primes.Sieve(limit)
	pm := primes.SieveMap(limit)
	for i := 4; i < limit; i++ {
		if !primes.IsPrime(i, pm) && !conjectureTrue(i, pl) {
			return i
		}
	}
	return 0
}
