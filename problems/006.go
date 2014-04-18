/*
Problem 6: Sum Square Difference
The sum of the squares of the first ten natural numbers is,
1² + 2² + ... + 10² = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)² = 55² = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package problems

const ANSWER_006 = 25164150

func square(n int) int {
	return n * n
}

func sumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func sumOfSeq(n int) int {
	return n * (n + 1) / 2
}

func Euler006() int {
	return square(sumOfSeq(100)) - sumOfSquares(100)
}

// func sumOfSquares(max int) int {
// 	var total int
// 	for i := 1; i <= max; i++ {
// 		total += square(i)
// 	}
// 	return total
// }

// func squareOfSum(max int) int {
// 	var total int
// 	for i := 1; i <= max; i++ {
// 		total += i
// 	}
// 	return square(total)
// }
