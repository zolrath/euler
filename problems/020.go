/*
Problem 20: Factorial Digit Sum
n! means n × (n − 1) × ... × 3 × 2 × 1
For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100!
Answer: 648
*/

package problems

import "math/big"

func factorialInt(num int64) *big.Int {
	n, one, result := big.NewInt(num), big.NewInt(1), big.NewInt(1)
	for i := big.NewInt(2); i.Cmp(n) != 1; i = i.Add(i, one) {
		result = result.Mul(result, i)
	}
	return result
}

// Uses sumDigits from 016.go
// TODO: Move reused/duplicated functions into util.

func Euler020() int {
	sum := sumDigits(factorialInt(100))
	return sum
}
