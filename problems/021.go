/*
Problem 21: Amicable Numbers
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
Evaluate the sum of all the amicable numbers under 10000.
Answer: 31626
*/

package problems

import (
	"runtime"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_021 = 31626

func isAmicable(a int) bool {
	b := sink.ProperDivisorSum(a)
	if a == b || b <= 0 {
		return false
	}
	s := sink.ProperDivisorSum(b)
	return a == s
}

func sumAmicable(n int) int {
	results := make(chan int)
	sum := 0

	for i := 1; i < n; i++ {
		// Parallelize it, heyo.
		go func(n int) {
			if isAmicable(n) {
				results <- n
			} else {
				results <- 0
			}
		}(i)
	}

	// Collect results.
	for i := 1; i < n; i++ {
		sum += <-results
	}

	return sum
}

func Euler021() int {
	runtime.GOMAXPROCS(runtime.NumCPU())
	return sumAmicable(10000)
}

// Prime Factorization to the rescue!
// σ(n) is the sum of divisors of a natural number.
// σ(p^a) = (p^a+1 − 1)/(p − 1).
// Get prime factorization of number,

// func properDivisorSum(n int) int {
// 	factors := primes.PrimeFactors(n)

// 	sum := 1
// 	for p, e := range factors {
// 		sum *= (int(math.Pow(float64(p), float64(e+1))) - 1) / (p - 1)
// 	}
// 	return sum - n
// }
