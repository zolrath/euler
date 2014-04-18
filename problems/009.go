/*
Problem 9: Special Pythagorean Triplet
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a² + b² = c²
For example, 3² + 4² = 9 + 16 = 25 = 5².
There exists exactly one Pythagorean triplet for which a + b + c = 1000.

Find the product abc.
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
