/*
Problem 47: Distinct Prime Factors
The first two consecutive numbers to have two distinct prime factors are:
14 = 2 × 7
15 = 3 × 5
The first three consecutive numbers to have three distinct prime factors are:
644 = 2² × 7  × 23
645 = 3  × 5  × 43
646 = 2  × 17 × 19

Find the first four consecutive integers to have four distinct prime factors.
What is the first of these numbers?
*/

package problems

import "github.com/zolrath/euler/util/primes"

const ANSWER_047 = 134043

func factorTotal(n int) (int, int) {
	total := 0
	factors := primes.PrimeFactors(n)
	for f, _ := range factors {
		total += f
	}
	return total, len(factors)
}

func Euler047() int {
	consec := 4
	for i := 2 * 3 * 5 * 7; i < 150000; i++ {
		lastFacTot := 0
		uniqConsec := true

		for j := 0; j < consec; j++ {
			facTot, length := factorTotal(i + j)

			if length != consec || lastFacTot == facTot {
				uniqConsec = false
				break
			}

			lastFacTot = facTot
		}

		if uniqConsec {
			return i
		}
	}
	return 0
}
