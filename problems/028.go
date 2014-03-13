/*
Problem 28: Number Spiral Diagonals
Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:
~21~ 22  23  24  ~25~
 20  ~7~  8  ~9~  10
 19   6  ~1~  2   11
 18  ~5~  4  ~3~  12
~17~ 16  15  14  ~13~
It can be verified that the sum of the numbers on the diagonals is 101.
What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
Answer: 669171001
*/

// Each spiral increases its width/height by 2. 1x1,3x3,5x5.
// Corner pieces can be found with (n*n)-n-1   n=3 for the 3x3 =  9     7      5      3
//                                                                    (9-2) (9-2*2) (9-2*3)

package problems

func diagonalCount(n int) int {
	total := 1
	for i := 3; i <= n; i += 2 {
		total += (i * i * 4) - (i-1)*6
	}
	return total
}

func Euler028() int {
	return diagonalCount(1001)
}
