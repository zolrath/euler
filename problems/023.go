/*
Problem 23: Non-Abdundant Sums
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28,
which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n
    and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16,
the smallest number that can be written as the sum of two abundant numbers is 24.
By mathematical analysis, it can be shown that all integers greater than 28123
can be written as the sum of two abundant numbers. However, this upper limit cannot
be reduced any further by analysis even though it is known that the greatest number
that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

package problems

import (
	"runtime"
	"sync"

	"github.com/zolrath/euler/util/sink"
)

const ANSWER_023 = 4179871

func isAbundant(n int) bool {
	return n < sink.ProperDivisorSum(n)
}

func Euler023() int {
	// Generate list of abundant numbers < 20162
	runtime.GOMAXPROCS(runtime.NumCPU())
	abundantList := []int{}
	for i := 1; i < 20162; i++ {
		if isAbundant(i) {
			abundantList = append(abundantList, i)
		}
	}

	// Generate all combinations of abundant sums.
	l := len(abundantList)
	abundantSums := [20162]bool{}

	// Lets try WaitGroups for this one.
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := n; j < l; j++ {
				sum := abundantList[n] + abundantList[j]
				if sum < 20162 {
					abundantSums[sum] = true
				}
			}
		}(i)
	}

	wg.Wait()

	// Find all numbers which aren't on the abundant sum list.
	total := 0
	for i := 1; i < 20162; i++ {
		if abundantSums[i] == false {
			total += i
		}
	}
	return total
}
