/*
Problem 51: Prime Digital Replacements
By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine possible values:
13, 23, 43, 53, 73, and 83, are all prime.
By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number is the first example
having seven primes among the ten generated numbers, yielding the family:
56003, 56113, 56333, 56443, 56663, 56773, and 56993.
Consequently 56003, being the first member of this family, is the smallest prime with this property.
Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits) with the same digit,
is part of an eight prime value family.
*/

package problems

import (
	"strconv"
	"strings"

	"github.com/zolrath/euler/util/primes"
)

const ANSWER_051 = 121313

func hasThreeDupes(n int) (string, string, bool) {
	ns := strconv.Itoa(n)
	for _, r := range ns {
		if strings.Count(ns, string(r)) == 3 {
			return ns, string(r), true
		}
	}
	return "", "", false
}

func Euler051() int {
	pm := primes.SieveMap(1000000)
	pl := primes.Sieve(1000000)
	for _, p := range pl {
		if ns, r, ok := hasThreeDupes(p); ok {
			pCount := 0
			for i := '0'; i <= '9'; i++ {
				pCandidate := strings.Replace(ns, r, string(i), -1)
				if pCandidate[0] == '0' {
					continue
				}
				pCandidateInt, _ := strconv.Atoi(pCandidate)

				if primes.IsPrime(pCandidateInt, pm) {
					pCount++
				}
			}
			if pCount == 8 {
				return p
			}
		}
	}
	return 0
}

// func replaceDigit(num, index, replacement int) int {
// 	digits := sink.SplitDigits(num)
// 	digits[index-1] = replacement
// 	return sink.JoinDigits(digits)
// }
