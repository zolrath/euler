/*
Problem 4: Largest Palindrome Product
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package problems

import "github.com/zolrath/euler/util/sink"

const ANSWER_004 = 906609

func Euler004() int {
	var largest int

	// By bounding j to less than i we don't perform repeated multiplication.
	for i := 100; i < 1000; i++ {
		for j := 100; j < i; j++ {
			result := i * j
			if sink.IsPalindromeInt(result) && result > largest {
				largest = result
			}
		}
	}
	return largest
}
