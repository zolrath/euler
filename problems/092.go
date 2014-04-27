/*
Problem 92: Square Digit Chains

A number chain is created by continuously adding the square of the digits in a number to form a new number until it has been seen before.
For example,

44 → 32 → 13 → 10 → 1 → 1
85 → 89 → 145 → 42 → 20 → 4 → 16 → 37 → 58 → 89

Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop.
What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.

How many starting numbers below ten million will arrive at 89?
*/

package problems

import "github.com/zolrath/euler/util/sink"

const ANSWER_092 = 8581146

func squareDigits(n int) int {
	digits := sink.SplitDigits(n)
	sum := 0
	for _, v := range digits {
		sum += v * v
	}
	return sum
}

func chainResult(n int, nc map[int]int) int {
	chain := make([]int, 0)
	for {
		if result, ok := nc[n]; ok {
			n = result
			break
		} else {
			chain = append(chain, n)
			n = squareDigits(n)
		}
	}
	for _, v := range chain {
		nc[v] = n
	}
	return n
}

func Euler092() int {
	sum := 0
	limit := 10000000
	nc := map[int]int{1: 1, 89: 89}
	for i := 2; i < limit; i++ {
		if chainResult(i, nc) == 89 {
			sum++
		}
	}
	return sum
}
