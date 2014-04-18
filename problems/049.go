/*
Problem 49: Prime Permutations
The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways:
(i) each of the three terms are prime, and,
(ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property,
but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?
*/

package problems

import (
	"github.com/zolrath/euler/util/permute"
	"github.com/zolrath/euler/util/primes"
	"github.com/zolrath/euler/util/sink"
)

const ANSWER_049 = 296962999629

func filterPrimes(nums []int, primeMap map[int]bool) []int {
	primeList := make([]int, 0)
	for _, v := range nums {
		if primes.IsPrime(v, primeMap) && sink.IntLen(v) == 4 {
			primeList = append(primeList, v)
		}
	}
	return primeList
}

func findIncreasing(nums []int) int {
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums); j++ {
			diff := nums[j] - nums[i]
			for k := j + i; k < len(nums); k++ {
				diff2 := nums[k] - nums[j]
				if diff2 == diff {
					s := sink.ConcatInt(nums[i], nums[j])
					s2 := sink.ConcatInt(s, nums[k])
					return s2
				}
			}
		}
	}
	return 0
}

func Euler049() int {
	pl := primes.Sieve(10000)
	pm := primes.SieveMap(100000)
	// Start at first 4 digit prime.
	for i := 168; i < len(pl); i++ {
		perms := permute.UniqInts(pl[i])
		if perms[0] == 1478 {
			continue
		}
		filtered := filterPrimes(perms, pm)
		result := findIncreasing(filtered)
		if result != 0 {
			return result
		}
	}
	return 0
}
