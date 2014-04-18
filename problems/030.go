/*
Problem 30: Digit Fifth Powers
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
1634 = 1⁴ + 6⁴ + 3⁴ + 4⁴
8208 = 8⁴ + 2⁴ + 0⁴ + 8⁴
9474 = 9⁴ + 4⁴ + 7⁴ + 4⁴
As 1 = 1⁴ is not a sum it is not included.
The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

package problems

import "math"

const ANSWER_030 = 443839

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
