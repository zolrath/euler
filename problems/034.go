/*
Problem 34: Digit Factorials
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.
NOTE: as 1! = 1 and 2! = 2 are not sums they are not included.
*/
package problems

import "github.com/zolrath/euler/util/sink"

const ANSWER_034 = 40730

func sumFactorials(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += sink.FactorialInt(v)
	}
	return sum
}

func Euler034() int {
	var sum int
	for i := int(3); i < 50000; i++ {
		digits := sink.SplitDigits(i)
		sumf := sumFactorials(digits)
		if sumf == i {
			sum += i
		}
	}
	return int(sum)
}

// func getDigits(n int) []int {
// 	digits := make([]int, 0)
// 	for n > 0 {
// 		dig := n % 10
// 		digits = append(digits, dig)
// 		n /= 10
// 	}
// 	return digits
// }
