/*
Problem 37: Truncatable Primes
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right,
and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.
Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/
package problems

import (
	"math"

	"github.com/zolrath/euler/util/primes"
	"github.com/zolrath/euler/util/sink"
)

const ANSWER_037 = 748317

func truncateLeft(n int) []int {
	nums := make([]int, 0)
	for iLen := sink.IntLen(n) - 1; iLen > 0; iLen-- {
		n %= int(math.Pow10(iLen))
		nums = append(nums, n)
	}
	return nums
}

func truncateRight(n int) []int {
	nums := make([]int, 0)
	for n > 10 {
		n /= 10
		nums = append(nums, n)
	}
	return nums
}

func allArePrime(nums []int, primeMap map[int]bool) bool {
	for _, p := range nums {
		if _, ok := primeMap[p]; !ok {
			return false
		}
	}
	return true
}

func Euler037() int {
	sum := 0
	max := 740000
	primeMap := primes.SieveMap(max)
	for p, _ := range primeMap {
		if p < 10 {
			continue
		}
		tl := truncateLeft(p)
		tr := truncateRight(p)
		if allArePrime(tl, primeMap) && allArePrime(tr, primeMap) {
			sum += p
		}
	}
	return sum
}
