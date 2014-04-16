/*
Problem 9: Special Pythagorean Triplet
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
Answer: 31875000
*/

package problems

const ANSWER_009 = 31875000

func findPyTrip(sum int) int {
	for a := 1; a < (sum)/3; a++ {
		for b := a + 1; b < (sum-a)/2; b++ {
			c := sum - a - b
			if c*c == a*a+b*b {
				return a * b * c
			}
		}
	}
	return -1
}

func Euler009() int {
	return findPyTrip(1000)
}
