/*
Problem 35: Circular Primes
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
How many circular primes are there below one million?
*/

package problems

import (
	"strconv"
	"strings"

	"github.com/zolrath/euler/util/sink"
	"github.com/zolrath/euler/util/primes"
)

const ANSWER_035 = 55

func allPrime(nums []int, primeMap map[int]bool) bool {
	for _, v := range nums {
		if !primes.IsPrime(v, primeMap) {
			return false
		}
	}
	return true
}

func Euler035() int {
	pm := primes.SieveMap(1000000)
	count := 3
	for i := 7; i < 1000000; i += 2 {
		ia := strconv.Itoa(i)
		if strings.ContainsAny(ia, "024568") {
			continue
		}
		perms := sink.GetIntRotations(i)
		if allPrime(perms, pm) {
			count += 1
		}
	}

	return count
}
