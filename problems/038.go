/*
Problem 38: Pandigital Multiples
Take the number 192 and multiply it by each of 1, 2, and 3:
192 × 1 = 192
192 × 2 = 384
192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576.

We will call 192384576 the concatenated product of 192 and (1,2,3)
The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5,
giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).
What is the largest 1 to 9 pandigital 9-digit number that can be formed as the
concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package problems

import (
	"math"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_038 = 932718654

func findPandigital(n int) int {
	lastPandigital := 0
	conj := 0
	for i := 1; i < 10 && sink.IntLen(conj) < 10; i++ {
		newDigits := n * i
		newLength := sink.IntLen(newDigits)
		conj *= int(math.Pow10(newLength))
		conj += newDigits

		if sink.IntLen(conj) == 9 && isPandigital(conj) {
			lastPandigital = conj
		}
	}
	return lastPandigital
}

func Euler038() int {
	largest := 0
	for i := 9999; i > 1000; i-- {
		pan := findPandigital(i)
		if pan > largest {
			return pan
		}
	}
	return largest
}
