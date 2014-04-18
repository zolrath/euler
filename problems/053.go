/*
Problem 53: Combinatoric Selections
There are exactly ten ways of selecting three from five, 12345:
123, 124, 125, 134, 135, 145, 234, 235, 245, and 345
In combinatorics, we use the notation, ⁵C₃ = 10.
In general,
          n!
ⁿCᵣ =  --------
       r!(n−r)!
where r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1.
It is not until n = 23, that a value exceeds one-million: ²³C₁₀ = 1144066.

How many, not necessarily distinct, values of ⁿCᵣ, for 1 ≤ n ≤ 100, are greater than one-million?
*/

package problems

import "github.com/zolrath/euler/util/sink"

const ANSWER_053 = 4075

func c(n, r float64) float64 {
	return sink.Factorial(n) / (sink.Factorial(r) * sink.Factorial(n-r))
}

func Euler053() int {
	sum := 0
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			if c(float64(n), float64(r)) > 1000000.0 {
				sum++
			}
		}
	}
	return sum
}
