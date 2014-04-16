/*
Problem 3: Largest Prime Factor
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143?
Answer: 6857
*/

package problems

const ANSWER_003 = 6857

func Euler003() int {
	var n, fac int
	n = 600851475143
	for n%2 == 0 {
		fac = 2
		n = n / fac
	}
	for i := 3; i <= n; i += 2 {
		for n%i == 0 {
			fac = i
			n = n / fac
		}
	}
	return fac
}
