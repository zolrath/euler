/*
Problem 36: Double-base Palindromes
The decimal number, 585 = 1001001001 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
NOTE: The palindromic number, in either base, may not include leading zeros.
*/
package problems

import (
	"strconv"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_036 = 872187

func isBinaryPalindrome(n int64) bool {
	binaryString := strconv.FormatInt(n, 2)
	return sink.IsPalindromeString(binaryString)
}

func Euler036() int {
	sum := 0
	for i := 1; i < 1000000; i++ {
		if sink.IsPalindromeInt(i) && isBinaryPalindrome(int64(i)) {
			sum += i
		}
	}
	return sum
}
