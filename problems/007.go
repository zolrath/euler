/*
Problem 7: 10001st Prime
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10,001st prime number?
*/

package problems

import "github.com/zolrath/euler/util/primes"

const ANSWER_007 = 104743

func Euler007() int {
	return primes.Nth(10001)
}
