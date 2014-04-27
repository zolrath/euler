/*
Problem 95: Amicable Chains
The proper divisors of a number are all the divisors excluding the number itself.
For example, the proper divisors of 28 are 1, 2, 4, 7, and 14.
As the sum of these divisors is equal to 28, we call it a perfect number.

Interestingly the sum of the proper divisors of 220 is 284 and the sum of the proper divisors of 284 is 220,
forming a chain of two numbers. For this reason, 220 and 284 are called an amicable pair.

Perhaps less well known are longer chains. For example, starting with 12496, we form a chain of five numbers:

12496 → 14288 → 15472 → 14536 → 14264 (→ 12496 → ...)

Since this chain returns to its starting point, it is called an amicable chain.

Find the smallest member of the longest amicable chain with no element exceeding one million.
*/

package problems

const ANSWER_095 = 14316

func amicableChain(start int, pdsCache []int) (int, int, bool) {
	members := map[int]bool{start: true}
	length := 1
	smallest := start
	isChain := false
	last := start

	for {
		pds := pdsCache[last]
		if pds >= 1000000 {
			break
		}
		// If we've already added this to this last we've entered a loop.
		if _, exists := members[pds]; exists {
			isChain = pds == start // if pds == start it is a chain.
			break
		}

		if pds < smallest {
			smallest = pds
		}

		members[pds] = true
		last = pds
		length++
	}
	return length, smallest, isChain
}

func generatePDS(limit int) []int {
	pdsCache := make([]int, limit)
	// Use a sieve to generate proper divisor sums.
	// At 2 find all numbers divisible by 2, add 2 to the sum. Move to 3. Repeat!
	for i := 1; i < (limit+1)/2; i++ {
		for j := i * 2; j < limit; j += i {
			pdsCache[j] += i
		}
	}
	return pdsCache
}

func Euler095() int {
	var longestChain, smallestMember int
	limit := 1000000
	pdsCache := generatePDS(limit)
	for i := 1; i < limit; i++ {
		if chainLength, smallest, isChain := amicableChain(i, pdsCache); isChain {
			if chainLength > longestChain {
				longestChain = chainLength
				smallestMember = smallest
			}
		}
	}
	return smallestMember
}

// func aliquotSequence(start int) []int {
// 	seq := []int{start}
// 	members := map[int]bool{start: true}
// 	last := start

// 	for {
// 		pds := sink.ProperDivisorSum(last)
// 		// If we've already added this to this last we've entered a loop.
// 		if _, exists := members[pds]; exists {
// 			break
// 		}

// 		seq = append(seq, pds)
// 		members[pds] = true
// 		last = pds
// 	}
// 	return seq
// }
