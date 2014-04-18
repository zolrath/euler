/*
Problem 50: Consecutive Prime Sums
The prime 41, can be written as the sum of six consecutive primes:
41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.
The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.
Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/

package problems

import "github.com/zolrath/euler/util/primes"

const ANSWER_050 = 997651

func Euler050() int {
	max := 1000000
	mostConsec := 0
	mostConsecPrime := 0

	pl := primes.Sieve(max)
	pm := primes.SieveMap(max)

	// Make longest sum < max for each prime starting number.
	for i := 0; i < len(pl)-1; i++ {
		consec := 1
		consecPrime := pl[i]
		for {
			consecPrime += pl[i+consec]
			consec++

			if consecPrime > max {
				break
			}

			if primes.IsPrime(consecPrime, pm) {
				if consec > mostConsec {
					mostConsec = consec
					mostConsecPrime = consecPrime
				}
			}
		}
	}
	return mostConsecPrime
}
