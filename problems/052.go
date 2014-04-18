/*
Problem: 52: Permuted Multiples
It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
*/

package problems

import (
	"sort"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_052 = 142857

func sameDigits(a, b int) bool {
	if sink.IntLen(a) != sink.IntLen(b) {
		return false
	}
	aDigits := sink.SplitDigits(a)
	bDigits := sink.SplitDigits(b)
	sort.Ints(aDigits)
	sort.Ints(bDigits)

	for i := 0; i < len(aDigits); i++ {
		if aDigits[i] != bDigits[i] {
			return false
		}
	}
	return true
}

func Euler052() int {
	for i := 100000; i < 1000000; i++ {
		allSame := true
		for m := 2; m <= 6; m++ {
			if !sameDigits(i, m*i) {
				allSame = false
				break
			}
		}
		if allSame {
			return i
		}
	}
	return 0
}
