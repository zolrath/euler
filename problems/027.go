/*
Problem 27: Quadratic Primes
Euler discovered the remarkable quadratic formula:
n² + n + 41
It turns out that the formula will produce 40 primes for the consecutive values
n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41,
and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes
for the consecutive values n = 0 to 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

n² + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n
e.g. |11| = 11 and |−4| = 4
Find the product of the coefficients, a and b, for the quadratic expression that
produces the maximum number of primes for consecutive values of n, starting with n = 0.
Answer: -59231
*/

package problems

import "github.com/zolrath/euler/util/primes"

const ANSWER_027 = -59231

var primesl map[int]bool

func quadForm(a, b int) int {
	for n := 0; ; n++ {
		if primes.IsPrime((n*n)+a*n+b, primesl) != true {
			return n
		}
	}
}

func Euler027() int {
	high, topa, topb := 0, 0, 0
	primesl = primes.SieveMap(1000)
	for a := -100; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			ps := quadForm(a, b)
			if ps > high {
				high = ps
				topa = a
				topb = b
			}
		}
	}
	return topa * topb
}
