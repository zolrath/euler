/*
Problem 30: Digit Fifth Powers
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
Answer: 443839
*/

package problems

import "math"

func sumPowDigits(n int) int {
	mod, sum := 0, 0
	for n > 0 {
		mod = n % 10
		n = n / 10
		sum += int(math.Pow(float64(mod), 5))
	}
	return sum
}

func Euler030() int {
	total := 0
	for i := 2; i < 354294; i++ {
		if sumPowDigits(i) == i {
			total += i
		}
	}
	return total
}
