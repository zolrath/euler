/*
Problem 3: Largest Prime Factor
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143?
Answer: 6857
*/

package problems

func Euler003() int {
	var n, fac int
	n = 600851475143
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			fac = i
			n = n / fac
		}
	}
	return fac
}
