/*
Problem 14: Longest Collatz Sequence
The following iterative sequence is defined for the set of positive integers:
	n → n/2 (n is even)
	n → 3n + 1 (n is odd)
Using the rule above and starting with 13, we generate the following sequence:
	13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?
NOTE: Once the chain starts the terms are allowed to go above one million.
Answer: 837799
*/

package problems

// Map for memoization
var collatzMap map[int]int

func collatz(n int) int {
	cycles := 1
	for i := n; i != 1; {
		if collatzMap[i] != 0 {
			cycles += collatzMap[i]
			break
		}
		switch i%2 == 0 {
		case true:
			i = i / 2
			cycles++
		case false:
			i = 3*i + 1
			cycles++
		}
	}
	collatzMap[n] = cycles
	return cycles
}

func Euler014() int {
	collatzMap = map[int]int{}
	longestChain := 0
	longestChainNum := 0
	for i := 1; i < 1000000; i++ {
		chain := collatz(i)
		if chain > longestChain {
			longestChain = chain
			longestChainNum = i
		}
	}
	return longestChainNum
}
