/*
Problem 6: Sum Square Difference
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
Answer: 25164150
*/

package problems

func square(n int) int {
	return n * n
}

func sumOfSquares(max int) int {
	var total int
	for i := 1; i <= max; i++ {
		total += square(i)
	}
	return total
}

func squareOfSum(max int) int {
	var total int
	for i := 1; i <= max; i++ {
		total += i
	}
	return square(total)
}

func Euler006() int {
	return squareOfSum(100) - sumOfSquares(100)
}
