/*
Problem 32: Pandigital Products
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once;
for example, the 5-digit number, 15234, is 1 through 5 pandigital.
The product 7254 is unusual, as the identity, 39 × 186 = 7254,
containing multiplicand, multiplier, and product is 1 through 9 pandigital.
Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
Answer: 45228
*/

package problems

import "github.com/zolrath/euler/util/sink"

const ANSWER_032 = 45228

func combine(a, b, c int) int {
	var total int
	for ; a > 0; a /= 10 {
		total = (total * 10) + (a % 10)
	}
	for ; b > 0; b /= 10 {
		total = (total * 10) + (b % 10)
	}
	for ; c > 0; c /= 10 {
		total = (total * 10) + (c % 10)
	}
	return total
}

func isPandigital(n int) bool {
	l := sink.IntLen(n)
	// For this problem we want all 9
	if l != 9 {
		return false
	}
	nums := [10]int{}
	for n > 0 {
		dig := n % 10
		if nums[dig] == 1 || dig == 0 {
			return false
		}
		nums[dig] = 1
		n /= 10
	}
	return true
}

func Euler032() int {
	pandigitals := map[int]bool{}
	for i := 1; i < 2000; i++ {
		for j := 1; j < 100; j++ {
			product := i * j
			if isPandigital(combine(i, j, product)) {
				pandigitals[product] = true
			}
		}
	}
	sum := 0
	for i, _ := range pandigitals {
		sum += i
	}
	return sum
}
