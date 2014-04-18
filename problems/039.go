/*
Problem 39: Integer Right Triangles
If p is the perimeter of a right angle triangle with integral length sides, {a,b,c},
there are exactly three solutions for p = 120.
{20,48,52}, {24,45,51}, {30,40,50}
For which value of p ≤ 1000, is the number of solutions maximised?
*/
package problems

const ANSWER_039 = 840

func findRightTriangles(p int) [][]int {
	tris := make([][]int, 0)
	for a := 1; a < (p)/3; a++ {
		for b := a + 1; b < (p-a)/2; b++ {
			c := p - a - b
			if c*c == a*a+b*b {
				tris = append(tris, []int{a, b, c})
			}
		}
	}
	return tris
}

func Euler039() int {
	solCount := 0
	solN := 0
	for p := 3; p < 1000; p++ {
		solutions := len(findRightTriangles(p))
		if solutions > solCount {
			solCount = solutions
			solN = p
		}
	}
	return solN
}
