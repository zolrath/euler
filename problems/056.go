/*
Problem 56: Powerful Digit Sum
A googol (10¹⁰⁰) is a massive number: one followed by one-hundred zeros;
100¹⁰⁰ is almost unimaginably large: one followed by two-hundred zeros.
Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, aᵇ, where a, b < 100, what is the maximum digital sum?
*/

package problems

import (
	"math/big"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_056 = 972

func Euler056() int {
	largest := 0
	for a := int64(1); a < 100; a++ {
		biga := big.NewInt(a)
		for b := int64(1); b < 100; b++ {
			bigb := big.NewInt(b)
			powed := big.NewInt(0).Exp(biga, bigb, nil)
			sum := sink.SumBigDigits(powed)
			if sum > largest {
				largest = sum
			}
		}
	}
	return largest
}
