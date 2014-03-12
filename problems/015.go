/*
Problem 15: Lattice Paths
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down,
there are exactly 6 routes to the bottom right corner.
How many such routes are there through a 20×20 grid?
Answer: 137846528820
*/

package problems

func factorial(n float64) float64 {
	result := float64(1)
	for i := float64(2); i <= n; i++ {
		result *= i
	}
	return result
}

// Combinatorics: (m+n)!/m!n!
func pathsToPoint(x, y float64) int {
	a := factorial(x + y)
	b := factorial(x)
	c := factorial(y)
	return int(a / (b * c))
}

func Euler015() int {
	return pathsToPoint(20, 20)
}
